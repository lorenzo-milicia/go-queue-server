package domain

type Record struct {
	ID     string
	Fields map[string]interface{}
}

type Records []Record

func (s BatchService) GetBatch(pagesize int, pagenumber int) []Record {
	return s.Repository.FindAllPaginated(pagesize, pagenumber)
}

func (s BatchService) GetBatchChannel(ch chan Records, pagesize int) {
	s.Repository.AsynchFetchRecords(ch, pagesize)
}
