// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: template.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/pipelinetemplate/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// TemplateService template.proto
	TemplateService() pb.TemplateServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		templateService: pb.NewTemplateServiceClient(cc),
	}
}

type serviceClients struct {
	templateService pb.TemplateServiceClient
}

func (c *serviceClients) TemplateService() pb.TemplateServiceClient {
	return c.templateService
}

type templateServiceWrapper struct {
	client pb.TemplateServiceClient
	opts   []grpc1.CallOption
}

func (s *templateServiceWrapper) ApplyPipelineTemplate(ctx context.Context, req *pb.PipelineTemplateApplyRequest) (*pb.PipelineTemplateCreateResponse, error) {
	return s.client.ApplyPipelineTemplate(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *templateServiceWrapper) QueryPipelineTemplates(ctx context.Context, req *pb.PipelineTemplateQueryRequest) (*pb.PipelineTemplateQueryResponse, error) {
	return s.client.QueryPipelineTemplates(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *templateServiceWrapper) RenderPipelineTemplate(ctx context.Context, req *pb.PipelineTemplateRenderRequest) (*pb.PipelineTemplateRenderResponse, error) {
	return s.client.RenderPipelineTemplate(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *templateServiceWrapper) RenderPipelineTemplateBySpec(ctx context.Context, req *pb.PipelineTemplateRenderSpecRequest) (*pb.PipelineTemplateRenderResponse, error) {
	return s.client.RenderPipelineTemplateBySpec(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *templateServiceWrapper) GetPipelineTemplateVersion(ctx context.Context, req *pb.PipelineTemplateVersionGetRequest) (*pb.PipelineTemplateVersionGetResponse, error) {
	return s.client.GetPipelineTemplateVersion(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *templateServiceWrapper) QueryPipelineTemplateVersions(ctx context.Context, req *pb.PipelineTemplateVersionQueryRequest) (*pb.PipelineTemplateVersionQueryResponse, error) {
	return s.client.QueryPipelineTemplateVersions(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *templateServiceWrapper) QuerySnippetYml(ctx context.Context, req *pb.QuerySnippetYmlRequest) (*pb.QuerySnippetYmlResponse, error) {
	return s.client.QuerySnippetYml(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
