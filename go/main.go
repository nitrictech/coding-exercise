package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nitrictech/coding-exercise/go/gen/interview/storage"
	"github.com/nitrictech/coding-exercise/go/server"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	storage.RegisterStorageServer(grpcServer, server.New())

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("unable to get port 50051")
	}

	fmt.Println("Listening on 50051")
	grpcServer.Serve(lis)
}
