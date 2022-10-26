package domain

type IBatchService interface {
	GetBatch(pagesize int, pagenumber int) []Record
	GetBatchChannel(ch chan Records, pagesize int)
}

type BatchService struct {
	Repository IRecordRepository
}

func (s BatchService) GetBatch(pagesize int, pagenumber int) []Record {
	return s.Repository.FindAllPaginated(pagesize, pagenumber)
}

func (s BatchService) GetBatchChannel(ch chan Records, pagesize int) {
	s.Repository.AsynchFetchRecords(ch, pagesize)
}