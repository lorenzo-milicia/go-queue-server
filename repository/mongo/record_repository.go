package repository

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoRecordRepository struct {
	DB *mongo.Client
}

func NewMongoRepository() *MongoRecordRepository {
	godotenv.Load(".env")
	// connect to hosted MongoDB

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("DB_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Ping to database failed")
	}
	log.Println("Ping successful")

	return &MongoRecordRepository{
		DB: client,
	}
}

// IRecordRepository implementation

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

func (r *MongoRecordRepository) AsynchFetchRecords(ch chan domain.Records, batchsize int) {
	log.Println("Start DB fetch")
	ctx := context.TODO()
	collection := r.DB.Database("go_queue_server").Collection("queue")

	opts := options.Find().
		SetSort(bson.D{{"_id", 1}})

	log.Println("Run find query...")

	cursor, err := collection.Find(
		ctx,
		bson.D{},
		opts,
	)

	log.Println("Cursor created")

	if err != nil {
		log.Fatalf("Error on Find statement: %v", err)
	}

	var batch []domain.Record = make([]domain.Record, 0)

	for cursor.Next(ctx) {

		var entity RecordEntity

		cursor.Decode(&entity)

		batch = append(batch, entity.toDomain())

		if len(batch) == batchsize {
			ch <- batch
			batch = make([]domain.Record, 0, batchsize)
		}
	}

	ch <- batch

	close(ch)
}
