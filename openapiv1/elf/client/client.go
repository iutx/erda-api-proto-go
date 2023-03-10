// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: elf.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/elf/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// Elf elf.proto
	Elf() pb.ElfClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		elf: pb.NewElfClient(cc),
	}
}

type serviceClients struct {
	elf pb.ElfClient
}

func (c *serviceClients) Elf() pb.ElfClient {
	return c.elf
}

type elfWrapper struct {
	client pb.ElfClient
	opts   []grpc1.CallOption
}

func (s *elfWrapper) ELF_ENVIROMENT_CREATE(ctx context.Context, req *pb.Environment) (*pb.EnvironmentResponse, error) {
	return s.client.ELF_ENVIROMENT_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_ENVIROMENT_DELETE(ctx context.Context, req *pb.ELF_ENVIROMENT_DELETE_Request) (*pb.EnvironmentResponse, error) {
	return s.client.ELF_ENVIROMENT_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_ENVIROMENT_GET(ctx context.Context, req *pb.ELF_ENVIROMENT_GET_Request) (*pb.EnvironmentResponse, error) {
	return s.client.ELF_ENVIROMENT_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_ENVIROMENT_LIST(ctx context.Context, req *pb.ELF_ENVIROMENT_LIST_Request) (*pb.EnvironmentListResponse, error) {
	return s.client.ELF_ENVIROMENT_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_ENVIROMENT_UPDATE(ctx context.Context, req *pb.Environment) (*pb.EnvironmentResponse, error) {
	return s.client.ELF_ENVIROMENT_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_NOTEBOOK_CREATE(ctx context.Context, req *pb.Notebook) (*pb.NotebookResponse, error) {
	return s.client.ELF_NOTEBOOK_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_NOTEBOOK_DELETE(ctx context.Context, req *pb.ELF_NOTEBOOK_DELETE_Request) (*pb.NotebookResponse, error) {
	return s.client.ELF_NOTEBOOK_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_NOTEBOOK_GET(ctx context.Context, req *pb.ELF_NOTEBOOK_GET_Request) (*pb.NotebookResponse, error) {
	return s.client.ELF_NOTEBOOK_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_NOTEBOOK_LIST(ctx context.Context, req *pb.ELF_NOTEBOOK_LIST_Request) (*pb.NotebookListResponse, error) {
	return s.client.ELF_NOTEBOOK_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_NOTEBOOK_UPDATE(ctx context.Context, req *pb.Notebook) (*pb.NotebookResponse, error) {
	return s.client.ELF_NOTEBOOK_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *elfWrapper) ELF_PACKAGE_LIST(ctx context.Context, req *pb.ELF_PACKAGE_LIST_Request) (*pb.DependencyPackageListResponse, error) {
	return s.client.ELF_PACKAGE_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
