package domain

type IConsumeQueueService interface {
	SaveBatches(ch chan Records) error
}

type ConsumeQueueService struct {
	Repository IRecordRepository
}

func (s ConsumeQueueService) SaveBatches(ch chan Records) error {
	for batch := range ch {
		err := s.Repository.SaveRecords(batch)
		if err != nil {
			return err
		}
	}
	return nil
}