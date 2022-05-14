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
	pagenumber := 0
	for {
		records := s.service.GetBatch(int(dto.Size), pagenumber)

		if len(records) == 0 {
			break
		}

		mappedRecords := api.Records{Records: toDto(records)}
		err := stream.Send(&mappedRecords)
		if err != nil {
			log.Fatalf("Error during stream send %v\n", err)
			return err
		} else {
			log.Print("Stream sent")
		}
		pagenumber++
	}
	return nil
}
