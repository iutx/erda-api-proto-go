// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: runtime.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/orchestrator/runtime/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// RuntimeService runtime.proto
	RuntimeService() pb.RuntimeServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		runtimeService: pb.NewRuntimeServiceClient(cc),
	}
}

type serviceClients struct {
	runtimeService pb.RuntimeServiceClient
}

func (c *serviceClients) RuntimeService() pb.RuntimeServiceClient {
	return c.runtimeService
}

type runtimeServiceWrapper struct {
	client pb.RuntimeServiceClient
	opts   []grpc1.CallOption
}

func (s *runtimeServiceWrapper) GetRuntime(ctx context.Context, req *pb.GetRuntimeRequest) (*pb.RuntimeInspect, error) {
	return s.client.GetRuntime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *runtimeServiceWrapper) CheckRuntimeExist(ctx context.Context, req *pb.CheckRuntimeExistReq) (*pb.CheckRuntimeExistResp, error) {
	return s.client.CheckRuntimeExist(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *runtimeServiceWrapper) DelRuntime(ctx context.Context, req *pb.DelRuntimeRequest) (*pb.Runtime, error) {
	return s.client.DelRuntime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
