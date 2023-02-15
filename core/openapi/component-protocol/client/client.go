// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: cp.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/openapi/component-protocol/pb"
	grpc1 "google.golang.org/grpc"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// Client provide all service clients.
type Client interface {
	// OpenapiComponentProtocol cp.proto
	OpenapiComponentProtocol() pb.OpenapiComponentProtocolClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		openapiComponentProtocol: pb.NewOpenapiComponentProtocolClient(cc),
	}
}

type serviceClients struct {
	openapiComponentProtocol pb.OpenapiComponentProtocolClient
}

func (c *serviceClients) OpenapiComponentProtocol() pb.OpenapiComponentProtocolClient {
	return c.openapiComponentProtocol
}

type openapiComponentProtocolWrapper struct {
	client pb.OpenapiComponentProtocolClient
	opts   []grpc1.CallOption
}

func (s *openapiComponentProtocolWrapper) Render(ctx context.Context, req *structpb.Value) (*structpb.Value, error) {
	return s.client.Render(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
