package domain

type IRecordRepository interface {
	FindAllPaginated(pagesize int, pagenumber int) []Record
	AsynchFetchRecords(ch chan Records, batchsize int)
	SaveRecords(records Records) error
}