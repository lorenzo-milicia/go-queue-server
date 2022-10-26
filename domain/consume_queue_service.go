package domain

type IConsumeQueueService interface {
	SaveBatches(records Records) error
}

type ConsumeQueueService struct {
	Repository IRecordRepository
}

func (s ConsumeQueueService) SaveBatches(records Records) error {
	err := s.Repository.SaveRecords(records)
	if err != nil {
		return err
	}
	return nil
}