package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Plugin adapter

var RecordRepository domain.IRecordRepository = newRepository()

func newRepository() *MongoRecordRepository {
	godotenv.Load(".env")
	// connect to hosted MongoDB

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("DB_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoRecordRepository{
		DB: client,
	}
}

//

type RecordEntity struct {
	ID     primitive.ObjectID     `bson:"_id"`
	Fields map[string]interface{} `bson:"data"`
}

type MongoRecordRepository struct {
	DB *mongo.Client
}

func (r *MongoRecordRepository) FindAllPaginated(pagesize int, pagenumber int) []domain.Record {
	ctx := context.TODO()
	db := r.DB.Database("go_queue_server")
	collection := db.Collection("queue")

	opts := options.Find().
		SetSort(bson.D{{"_id", 1}}).
		SetSkip(int64(pagenumber) * int64(pagesize)).
		SetLimit(int64(pagesize))

	cursor, err := collection.Find(
		ctx,
		bson.D{},
		opts,
	)

	if err != nil {
		log.Fatalf("Error on Find statement: %v", err)
	}

	var entities = make([]RecordEntity, 0)

	if err := cursor.All(ctx, &entities); err != nil {
		log.Fatalf("Error on All: %v", err)
	}

	var records = make([]domain.Record, 0)

	for _, e := range entities {
		records = append(records, e.toDomain())
	}

	return records
}
