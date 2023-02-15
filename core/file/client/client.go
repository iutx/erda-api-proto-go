// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: file.proto

package client

import (
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/file/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// FileService file.proto
	FileService() pb.FileServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		fileService: pb.NewFileServiceClient(cc),
	}
}

type serviceClients struct {
	fileService pb.FileServiceClient
}

func (c *serviceClients) FileService() pb.FileServiceClient {
	return c.fileService
}

type fileServiceWrapper struct {
	client pb.FileServiceClient
	opts   []grpc1.CallOption
}
