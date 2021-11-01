package main

import (
	"context"
	"log"
	"time"

	"github.com/nitrictech/coding-exercise/go/gen/interview/storage"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*5),
	)

	if err != nil {
		log.Fatalf("error connecting to storage server: %s", err.Error())
	}

	defer conn.Close()

	c := storage.NewStorageClient(conn)

	r, err := c.Read(context.TODO(), &storage.ReadRequest{
		Key: "test.txt",
	})

	if err != nil {
		log.Fatalf("error calling read: %s", err.Error())
	}

	log.Printf("content: %s", r.Content)
}
