package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path"
	"plugin"

	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/api_impl"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"google.golang.org/grpc"
)

func main() {

	// Define flags
	var repoModFile string
	flag.StringVar(&repoModFile, "repository", "", "The path to the repository plugin .so file")

	// Read flags
	flag.Parse()

	if repoModFile == "" {
		panic("No repository plugin specified\n")
	}

	//

	fmt.Printf("Opening plugin...\n")

	mod, err := plugin.Open(repoModFile)

	if err != nil {
		panic(fmt.Sprintf("Error while opening plugin file: %v\n", err))
	}

	sym, err := mod.Lookup("RecordRepository")

	if err != nil {
		panic(fmt.Sprintf("Error while looking up the variable: %v\n", err))
	}

	r := sym.(*domain.IRecordRepository)

	fmt.Printf("Repository plugin %v correctly loaded\n", path.Base(repoModFile))

	s := domain.BatchService{Repository: *r}

	fmt.Println("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDataFetcherServer(grpcServer, api_impl.NewDataFetcherServer(&s))
	grpcServer.Serve(lis)
}
