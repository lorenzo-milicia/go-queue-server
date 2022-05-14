package domain

type Record struct {
	ID     string
	Fields map[string]interface{}
}

func (s BatchService) GetBatch(pagesize int, pagenumber int) []Record {
	return s.Repository.FindAllPaginated(pagesize, pagenumber)
}
