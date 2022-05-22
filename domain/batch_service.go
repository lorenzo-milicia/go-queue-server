package domain

type IBatchService interface {
	GetBatch(pagesize int, pagenumber int) []Record
	GetBatchChannel(ch chan Records, pagesize int)
}

type IRecordRepository interface {
	FindAllPaginated(pagesize int, pagenumber int) []Record
	AsynchFetchRecords(ch chan Records, batchsize int)
}

type BatchService struct {
	Repository IRecordRepository
}
