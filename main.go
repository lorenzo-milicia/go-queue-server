package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/api_impl"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"github.com/lorenzo-milicia/go-server-queue/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting...")
	godotenv.Load(".env")
	// connect to hosted MongoDB

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("DB_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	r := repository.MongoRecordRepository{DB: client}

	s := domain.BatchService{Repository: &r}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDataFetcherServer(grpcServer, api_impl.NewDataFetcherServer(&s))
	grpcServer.Serve(lis)
}
