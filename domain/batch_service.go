package domain

type IBatchService interface {
	GetBatch(pagesize int, pagenumber int) []Record
}

type IRecordRepository interface {
	FindAllPaginated(pagesize int, pagenumber int) []Record
}

type BatchService struct {
	Repository IRecordRepository
}
