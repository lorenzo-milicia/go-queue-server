package main

import (
	"github.com/lorenzo-milicia/go-server-queue/domain"
)

// Plugin adapter

var RecordRepository domain.IRecordRepository = newRepository()

//

func newRepository() *MockRecordRepository {
	return &MockRecordRepository{}
}

type MockRecordRepository struct{}

func (r *MockRecordRepository) FindAllPaginated(pagesize int, pagenumber int) []domain.Record {
	if pagenumber < 10 {
		return make([]domain.Record, pagesize)
	} else {
		return make([]domain.Record, 0)
	}
}

func (r *MockRecordRepository) AsynchFetchRecords(ch chan domain.Records, batchsize int)
