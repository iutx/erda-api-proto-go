// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: core.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/issue/core/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// IssueCoreService core.proto
	IssueCoreService() pb.IssueCoreServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		issueCoreService: pb.NewIssueCoreServiceClient(cc),
	}
}

type serviceClients struct {
	issueCoreService pb.IssueCoreServiceClient
}

func (c *serviceClients) IssueCoreService() pb.IssueCoreServiceClient {
	return c.issueCoreService
}

type issueCoreServiceWrapper struct {
	client pb.IssueCoreServiceClient
	opts   []grpc1.CallOption
}

func (s *issueCoreServiceWrapper) CreateIssue(ctx context.Context, req *pb.IssueCreateRequest) (*pb.IssueCreateResponse, error) {
	return s.client.CreateIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) PagingIssue(ctx context.Context, req *pb.PagingIssueRequest) (*pb.PagingIssueResponse, error) {
	return s.client.PagingIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return s.client.GetIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UpdateIssue(ctx context.Context, req *pb.UpdateIssueRequest) (*pb.UpdateIssueResponse, error) {
	return s.client.UpdateIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) DeleteIssue(ctx context.Context, req *pb.DeleteIssueRequest) (*pb.DeleteIssueResponse, error) {
	return s.client.DeleteIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) BatchUpdateIssue(ctx context.Context, req *pb.BatchUpdateIssueRequest) (*pb.BatchUpdateIssueResponse, error) {
	return s.client.BatchUpdateIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UpdateIssueType(ctx context.Context, req *pb.UpdateIssueTypeRequest) (*pb.UpdateIssueTypeResponse, error) {
	return s.client.UpdateIssueType(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) SubscribeIssue(ctx context.Context, req *pb.SubscribeIssueRequest) (*pb.SubscribeIssueResponse, error) {
	return s.client.SubscribeIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UnsubscribeIssue(ctx context.Context, req *pb.UnsubscribeIssueRequest) (*pb.UnsubscribeIssueResponse, error) {
	return s.client.UnsubscribeIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) BatchUpdateIssueSubscriber(ctx context.Context, req *pb.BatchUpdateIssueSubscriberRequest) (*pb.BatchUpdateIssueSubscriberResponse, error) {
	return s.client.BatchUpdateIssueSubscriber(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) CreateIssueProperty(ctx context.Context, req *pb.CreateIssuePropertyRequest) (*pb.CreateIssuePropertyResponse, error) {
	return s.client.CreateIssueProperty(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) DeleteIssueProperty(ctx context.Context, req *pb.DeleteIssuePropertyRequest) (*pb.DeleteIssuePropertyResponse, error) {
	return s.client.DeleteIssueProperty(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UpdateIssueProperty(ctx context.Context, req *pb.UpdateIssuePropertyRequest) (*pb.UpdateIssuePropertyResponse, error) {
	return s.client.UpdateIssueProperty(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssueProperty(ctx context.Context, req *pb.GetIssuePropertyRequest) (*pb.GetIssuePropertyResponse, error) {
	return s.client.GetIssueProperty(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UpdateIssuePropertiesIndex(ctx context.Context, req *pb.UpdateIssuePropertiesIndexRequest) (*pb.UpdateIssuePropertiesIndexResponse, error) {
	return s.client.UpdateIssuePropertiesIndex(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssuePropertyUpdateTime(ctx context.Context, req *pb.GetIssuePropertyUpdateTimeRequest) (*pb.GetIssuePropertyUpdateTimeResponse, error) {
	return s.client.GetIssuePropertyUpdateTime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) CreateIssuePropertyInstance(ctx context.Context, req *pb.CreateIssuePropertyInstanceRequest) (*pb.CreateIssuePropertyInstanceResponse, error) {
	return s.client.CreateIssuePropertyInstance(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssuePropertyInstance(ctx context.Context, req *pb.GetIssuePropertyInstanceRequest) (*pb.GetIssuePropertyInstanceResponse, error) {
	return s.client.GetIssuePropertyInstance(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssueStage(ctx context.Context, req *pb.IssueStageRequest) (*pb.GetIssueStageResponse, error) {
	return s.client.GetIssueStage(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UpdateIssueStage(ctx context.Context, req *pb.IssueStageRequest) (*pb.UpdateIssueStageResponse, error) {
	return s.client.UpdateIssueStage(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) AddIssueRelation(ctx context.Context, req *pb.AddIssueRelationRequest) (*pb.AddIssueRelationResponse, error) {
	return s.client.AddIssueRelation(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) DeleteIssueRelation(ctx context.Context, req *pb.DeleteIssueRelationRequest) (*pb.DeleteIssueRelationResponse, error) {
	return s.client.DeleteIssueRelation(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssueRelations(ctx context.Context, req *pb.GetIssueRelationsRequest) (*pb.GetIssueRelationsResponse, error) {
	return s.client.GetIssueRelations(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) CreateIssueState(ctx context.Context, req *pb.CreateIssueStateRequest) (*pb.CreateIssueStateResponse, error) {
	return s.client.CreateIssueState(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) DeleteIssueState(ctx context.Context, req *pb.DeleteIssueStateRequest) (*pb.DeleteIssueStateResponse, error) {
	return s.client.DeleteIssueState(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) UpdateIssueStateRelation(ctx context.Context, req *pb.UpdateIssueStateRelationRequest) (*pb.UpdateIssueStateRelationResponse, error) {
	return s.client.UpdateIssueStateRelation(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssueStates(ctx context.Context, req *pb.GetIssueStatesRequest) (*pb.GetIssueStatesResponse, error) {
	return s.client.GetIssueStates(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) GetIssueStateRelation(ctx context.Context, req *pb.GetIssueStateRelationRequest) (*pb.GetIssueStateRelationResponse, error) {
	return s.client.GetIssueStateRelation(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) ExportExcelIssue(ctx context.Context, req *pb.ExportExcelIssueRequest) (*pb.ExportExcelIssueResponse, error) {
	return s.client.ExportExcelIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *issueCoreServiceWrapper) ImportExcelIssue(ctx context.Context, req *pb.ImportExcelIssueRequest) (*pb.ImportExcelIssueResponse, error) {
	return s.client.ImportExcelIssue(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}