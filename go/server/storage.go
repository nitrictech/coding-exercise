package server

import (
	"context"

	"github.com/nitrictech/coding-exercise/go/gen/interview/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorageServer struct {
	storage.UnimplementedStorageServer
}

func (*StorageServer) Read(ctx context.Context, req *storage.ReadRequest) (*storage.ReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "UNIMPLEMENTED")
}

func New() *StorageServer {
	return &StorageServer{}
}
