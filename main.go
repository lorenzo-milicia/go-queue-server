package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/api_impl"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	repository "github.com/lorenzo-milicia/go-server-queue/repository/mongo"
	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	r := repository.NewMongoRepository()

	s := domain.BatchService{Repository: r}

	fmt.Println("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDataFetcherServer(grpcServer, api_impl.NewDataFetcherServer(&s))
	_ = grpcServer.Serve(lis)
}
