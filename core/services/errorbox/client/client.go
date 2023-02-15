// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: errorbox.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb1 "github.com/erda-project/erda-proto-go/core/dop/taskerror/pb"
	pb "github.com/erda-project/erda-proto-go/core/services/errorbox/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ErrorBoxService errorbox.proto
	ErrorBoxService() pb.ErrorBoxServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		errorBoxService: pb.NewErrorBoxServiceClient(cc),
	}
}

type serviceClients struct {
	errorBoxService pb.ErrorBoxServiceClient
}

func (c *serviceClients) ErrorBoxService() pb.ErrorBoxServiceClient {
	return c.errorBoxService
}

type errorBoxServiceWrapper struct {
	client pb.ErrorBoxServiceClient
	opts   []grpc1.CallOption
}

func (s *errorBoxServiceWrapper) ListErrorLog(ctx context.Context, req *pb.TaskErrorListRequest) (*pb1.ErrorLogListResponseData, error) {
	return s.client.ListErrorLog(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
