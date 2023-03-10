// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: graph.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/graph/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// GraphService graph.proto
	GraphService() pb.GraphServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		graphService: pb.NewGraphServiceClient(cc),
	}
}

type serviceClients struct {
	graphService pb.GraphServiceClient
}

func (c *serviceClients) GraphService() pb.GraphServiceClient {
	return c.graphService
}

type graphServiceWrapper struct {
	client pb.GraphServiceClient
	opts   []grpc1.CallOption
}

func (s *graphServiceWrapper) PipelineYmlGraph(ctx context.Context, req *pb.PipelineYmlGraphRequest) (*pb.PipelineYmlGraphResponse, error) {
	return s.client.PipelineYmlGraph(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
