package api_impl

import (
	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/domain"
)

type QueueConsumerServerImpl struct {
	api.UnimplementedQueueConsumerServer
	service domain.IConsumeQueueService
}

func NewQueueConsumerServerImpl(s domain.IConsumeQueueService) *QueueConsumerServerImpl {
	return &QueueConsumerServerImpl{service: s}
}

func (s QueueConsumerServerImpl) ConsumeQueue(api.QueueConsumer_ConsumeQueueServer) error {
	panic("TODO")
}