// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: testplatform_testplan.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/testplatform/testplan/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// TestplatformTestplan testplatform_testplan.proto
	TestplatformTestplan() pb.TestplatformTestplanClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		testplatformTestplan: pb.NewTestplatformTestplanClient(cc),
	}
}

type serviceClients struct {
	testplatformTestplan pb.TestplatformTestplanClient
}

func (c *serviceClients) TestplatformTestplan() pb.TestplatformTestplanClient {
	return c.testplatformTestplan
}

type testplatformTestplanWrapper struct {
	client pb.TestplatformTestplanClient
	opts   []grpc1.CallOption
}

func (s *testplatformTestplanWrapper) AUTOTESTS_TESTPLAN_EXECUTE(ctx context.Context, req *pb.AUTOTESTS_TESTPLAN_EXECUTE_Request) (*emptypb.Empty, error) {
	return s.client.AUTOTESTS_TESTPLAN_EXECUTE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) AUTOTESTS_TESTPLAN_PAGING(ctx context.Context, req *pb.AUTOTESTS_TESTPLAN_PAGING_Request) (*emptypb.Empty, error) {
	return s.client.AUTOTESTS_TESTPLAN_PAGING(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) CANCEL_APITEST(ctx context.Context, req *pb.CANCEL_APITEST_Request) (*emptypb.Empty, error) {
	return s.client.CANCEL_APITEST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) CREATE(ctx context.Context, req *pb.TestPlanCreateRequest) (*pb.TestPlanCreateResponse, error) {
	return s.client.CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) DELETE(ctx context.Context, req *pb.DELETE_Request) (*emptypb.Empty, error) {
	return s.client.DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) EXECUTE_APITEST(ctx context.Context, req *pb.EXECUTE_APITEST_Request) (*emptypb.Empty, error) {
	return s.client.EXECUTE_APITEST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) GENERATE_REPORT(ctx context.Context, req *pb.GENERATE_REPORT_Request) (*pb.TestPlanReportGenerateResponse, error) {
	return s.client.GENERATE_REPORT(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) GET(ctx context.Context, req *pb.GET_Request) (*pb.TestPlanGetResponse, error) {
	return s.client.GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) LIST(ctx context.Context, req *pb.TestPlanPagingRequest) (*pb.TestPlanPagingResponse, error) {
	return s.client.LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) LIST_TESTSETS(ctx context.Context, req *pb.TestPlanTestSetsListRequest) (*pb.TestPlanTestSetListResponse, error) {
	return s.client.LIST_TESTSETS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *testplatformTestplanWrapper) UPDATE(ctx context.Context, req *pb.TestPlanUpdateRequest) (*emptypb.Empty, error) {
	return s.client.UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}