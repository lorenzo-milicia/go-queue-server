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
			err := stream.SendAndClose(&emptypb.Empty{})
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			log.Print("ERROR RECEIVING: ", err)
			return err
		}
		numberOfRecords := len(recv.Records)
		var mappedRecords = make([]*domain.Record, numberOfRecords)


		for i, record := range recv.Records {
			mappedRecords[i] = toDomain(record)
		}
		err = s.service.SaveBatches(mappedRecords)
		if err != nil {
			return err
		}
		log.Print("Message processed")
	}
}