package api_impl

import (
	"log"

	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/domain"
)

type DataFetcherServerImpl struct {
	api.UnimplementedDataFetcherServer
	service domain.IBatchService
}
func NewDataFetcherServer(s domain.IBatchService) *DataFetcherServerImpl {
	return &DataFetcherServerImpl{service: s}
}

func (s DataFetcherServerImpl) FetchQueueStream(dto *api.StreamSize, stream api.DataFetcher_FetchQueueStreamServer) error {
	log.Println("Start stream")

	ch := make(chan domain.Records, 2)

	go s.service.GetBatchChannel(ch, int(dto.Size))

	for batch := range ch {
		var mappedRecords = make([]*api.Record, 0)

		for _, record := range batch {
			mappedRecords = append(mappedRecords, toDto(record))
		}

		err := stream.Send(&api.Records{Records: mappedRecords})
		if err != nil {
			log.Fatalf("Error during stream send %v\n", err)
			return err
		} else {
			log.Print("Stream sent")
		}
	}
	return nil
}