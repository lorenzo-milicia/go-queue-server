package main

import (
	"fmt"
	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/api_impl"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	repository "github.com/lorenzo-milicia/go-server-queue/repository/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

//	_ = os.Setenv("DB_URL", "postgres://192.168.1.17:50001/queueserver")
//	_ = os.Setenv("DB_USERNAME", "queueserver")
//	_ = os.Setenv("DB_PASSWORD", "queueserver")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	r := repository.NewPSQLRecordRepository() 

//	batchService := domain.BatchService{Repository: r}
	consumeQueueService := domain.ConsumeQueueService{Repository: r}

	fmt.Println("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
//	api.RegisterDataFetcherServer(grpcServer, api_impl.NewDataFetcherServer(&batchService))
	api.RegisterQueueConsumerServer(grpcServer, api_impl.NewQueueConsumerServerImpl(&consumeQueueService))
	_ = grpcServer.Serve(lis)
}
