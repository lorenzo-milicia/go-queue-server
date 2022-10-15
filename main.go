package main

import (
	"fmt"
	repository "github.com/lorenzo-milicia/go-server-queue/repository/postgres"
	"log"
	"net"

	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/api_impl"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	r := repository.NewPSQLRecordRepository() 

	batchService := domain.BatchService{Repository: r}
	consumeQueueService := domain.ConsumeQueueService{Repository: r}

	fmt.Println("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDataFetcherServer(grpcServer, api_impl.NewDataFetcherServer(&batchService))
	api.RegisterQueueConsumerServer(grpcServer, api_impl.NewQueueConsumerServerImpl(&consumeQueueService))
	_ = grpcServer.Serve(lis)
}
