package api_impl

import (
	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

type QueueConsumerServerImpl struct {
	api.UnimplementedQueueConsumerServer
	service domain.IConsumeQueueService
}

func NewQueueConsumerServerImpl(s domain.IConsumeQueueService) *QueueConsumerServerImpl {
	return &QueueConsumerServerImpl{service: s}
}

func (s QueueConsumerServerImpl) ConsumeQueue(stream api.QueueConsumer_ConsumeQueueServer) error {
	log.Print("Started to receive messages...")
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Print("EOF")
			return stream.SendAndClose(&emptypb.Empty{})
		}
		var mappedRecords = make([]*domain.Record, 0)

		for _, record := range recv.Records {
			mappedRecords = append(mappedRecords, toDomain(record))
		}
		err = s.service.SaveBatches(mappedRecords)
		if err != nil {
			return err
		}
		log.Print("Message processed")
	}
}