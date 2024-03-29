package repository

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"log"
	"os"
)

type PSQLRecordRepository struct {
	DB *pgxpool.Pool
}

func NewPSQLRecordRepository() *PSQLRecordRepository {
	dbUrl := os.Getenv("DB_URL") + "?user=" + os.Getenv("DB_USERNAME") + "&password=" + os.Getenv("DB_PASSWORD") + "&pool_max_conns=10"
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Error while getting connection pool")
	}

	err = dbpool.Ping(context.TODO())
	if err != nil {
		log.Fatal("Failed ping", err)
	}
	log.Println("Ping successful")

	initDb(dbpool)
	log.Println("DB initialized correctly")

	return &PSQLRecordRepository{DB: dbpool}
}

func initDb(db *pgxpool.Pool) {
	file, err := os.ReadFile("./resources/sql/initDb.sql")
	if err != nil {
		log.Fatal("Unable to open initDb.sql")
	}

	_, err = db.Exec(context.TODO(), string(file))
	if err != nil {
		log.Fatal("Failed to initialize DB")
	}
}

// IRecordRepository implementation

func (r *PSQLRecordRepository) FindAllPaginated(pagesize int, pagenumber int) []domain.Record {
	panic("TODO")
}

func (r *PSQLRecordRepository) AsynchFetchRecords(ch chan domain.Records, batchsize int) {
	panic("TODO")
}

func (r *PSQLRecordRepository) SaveRecords(records domain.Records) error {
	batch := &pgx.Batch{}
	for _, record := range records {
		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		batch.Queue(
			"insert into data(id, payload) values($1, $2)",
			id.String(),
			record.Payload,
		)
	}
	br := r.DB.SendBatch(context.Background(), batch)
	err := br.Close()
	if err != nil {
		log.Print("Something went wrong with batch insert", err)
		return err
	}
	return nil
}
