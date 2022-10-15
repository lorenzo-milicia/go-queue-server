package repository

import (
	"context"
	_ "github.com/jackc/pgx/v4"
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
	defer dbpool.Close()

	err = dbpool.Ping(context.TODO())
	if err != nil {
		log.Fatal("Failed ping", err)
	}

	return &PSQLRecordRepository{DB: dbpool}
}

// IRecordRepository implementation

func (r *PSQLRecordRepository) FindAllPaginated(pagesize int, pagenumber int) []domain.Record {
	panic("TODO")
}

func (r *PSQLRecordRepository) AsynchFetchRecords(ch chan domain.Records, batchsize int)  {
	panic("TODO")
}

func (r *PSQLRecordRepository) SaveRecords(records domain.Records) error {
	panic("TODO")
}

