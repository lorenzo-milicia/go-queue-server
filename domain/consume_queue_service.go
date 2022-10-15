package domain

type IConsumeQueueService interface {
	SaveBatches(ch chan Records) error
}

type ConsumeQueueService struct {
	Repository IRecordRepository
}

func (s ConsumeQueueService) SaveBatches(ch chan Records) error {
	panic("TODO")
}