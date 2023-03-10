// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: testplatform_testplan_rel.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// TestplatformTestplanRelClient is the client API for TestplatformTestplanRel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestplatformTestplanRelClient interface {
	ADD_ISSUE_RELATION(ctx context.Context, in *TestPlanCaseRelIssueRelationAddRequest, opts ...grpc.CallOption) (*TestPlanCaseRelIssueRelationAddResponse, error)
	BATCH_UPDATE(ctx context.Context, in *TestPlanCaseRelBatchUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CREATE(ctx context.Context, in *TestPlanCaseRelCreateRequest, opts ...grpc.CallOption) (*TestPlanCaseRelCreateResponse, error)
	EXPORT(ctx context.Context, in *TestPlanCaseRelExportRequest, opts ...grpc.CallOption) (*TestPlanCaseRelExportResponse, error)
	GET(ctx context.Context, in *TestPlanCaseRelGetRequest, opts ...grpc.CallOption) (*TestPlanCaseRelGetResponse, error)
	PAGING(ctx context.Context, in *TestPlanCaseRelPagingRequest, opts ...grpc.CallOption) (*TestPlanCaseRelPagingResponse, error)
	REMOVE_ISSUE_RELATION(ctx context.Context, in *TestPlanCaseRelIssueRelationRemoveRequest, opts ...grpc.CallOption) (*TestPlanCaseRelIssueRelationRemoveResponse, error)
}

type testplatformTestplanRelClient struct {
	cc grpc1.ClientConnInterface
}

func NewTestplatformTestplanRelClient(cc grpc1.ClientConnInterface) TestplatformTestplanRelClient {
	return &testplatformTestplanRelClient{cc}
}

func (c *testplatformTestplanRelClient) ADD_ISSUE_RELATION(ctx context.Context, in *TestPlanCaseRelIssueRelationAddRequest, opts ...grpc.CallOption) (*TestPlanCaseRelIssueRelationAddResponse, error) {
	out := new(TestPlanCaseRelIssueRelationAddResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/ADD_ISSUE_RELATION", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testplatformTestplanRelClient) BATCH_UPDATE(ctx context.Context, in *TestPlanCaseRelBatchUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/BATCH_UPDATE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testplatformTestplanRelClient) CREATE(ctx context.Context, in *TestPlanCaseRelCreateRequest, opts ...grpc.CallOption) (*TestPlanCaseRelCreateResponse, error) {
	out := new(TestPlanCaseRelCreateResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/CREATE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testplatformTestplanRelClient) EXPORT(ctx context.Context, in *TestPlanCaseRelExportRequest, opts ...grpc.CallOption) (*TestPlanCaseRelExportResponse, error) {
	out := new(TestPlanCaseRelExportResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/EXPORT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testplatformTestplanRelClient) GET(ctx context.Context, in *TestPlanCaseRelGetRequest, opts ...grpc.CallOption) (*TestPlanCaseRelGetResponse, error) {
	out := new(TestPlanCaseRelGetResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/GET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testplatformTestplanRelClient) PAGING(ctx context.Context, in *TestPlanCaseRelPagingRequest, opts ...grpc.CallOption) (*TestPlanCaseRelPagingResponse, error) {
	out := new(TestPlanCaseRelPagingResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/PAGING", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testplatformTestplanRelClient) REMOVE_ISSUE_RELATION(ctx context.Context, in *TestPlanCaseRelIssueRelationRemoveRequest, opts ...grpc.CallOption) (*TestPlanCaseRelIssueRelationRemoveResponse, error) {
	out := new(TestPlanCaseRelIssueRelationRemoveResponse)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.testplatform.testplatform_testplan_rel/REMOVE_ISSUE_RELATION", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestplatformTestplanRelServer is the server API for TestplatformTestplanRel service.
// All implementations should embed UnimplementedTestplatformTestplanRelServer
// for forward compatibility
type TestplatformTestplanRelServer interface {
	ADD_ISSUE_RELATION(context.Context, *TestPlanCaseRelIssueRelationAddRequest) (*TestPlanCaseRelIssueRelationAddResponse, error)
	BATCH_UPDATE(context.Context, *TestPlanCaseRelBatchUpdateRequest) (*emptypb.Empty, error)
	CREATE(context.Context, *TestPlanCaseRelCreateRequest) (*TestPlanCaseRelCreateResponse, error)
	EXPORT(context.Context, *TestPlanCaseRelExportRequest) (*TestPlanCaseRelExportResponse, error)
	GET(context.Context, *TestPlanCaseRelGetRequest) (*TestPlanCaseRelGetResponse, error)
	PAGING(context.Context, *TestPlanCaseRelPagingRequest) (*TestPlanCaseRelPagingResponse, error)
	REMOVE_ISSUE_RELATION(context.Context, *TestPlanCaseRelIssueRelationRemoveRequest) (*TestPlanCaseRelIssueRelationRemoveResponse, error)
}

// UnimplementedTestplatformTestplanRelServer should be embedded to have forward compatible implementations.
type UnimplementedTestplatformTestplanRelServer struct {
}

func (*UnimplementedTestplatformTestplanRelServer) ADD_ISSUE_RELATION(context.Context, *TestPlanCaseRelIssueRelationAddRequest) (*TestPlanCaseRelIssueRelationAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ADD_ISSUE_RELATION not implemented")
}
func (*UnimplementedTestplatformTestplanRelServer) BATCH_UPDATE(context.Context, *TestPlanCaseRelBatchUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BATCH_UPDATE not implemented")
}
func (*UnimplementedTestplatformTestplanRelServer) CREATE(context.Context, *TestPlanCaseRelCreateRequest) (*TestPlanCaseRelCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CREATE not implemented")
}
func (*UnimplementedTestplatformTestplanRelServer) EXPORT(context.Context, *TestPlanCaseRelExportRequest) (*TestPlanCaseRelExportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EXPORT not implemented")
}
func (*UnimplementedTestplatformTestplanRelServer) GET(context.Context, *TestPlanCaseRelGetRequest) (*TestPlanCaseRelGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GET not implemented")
}
func (*UnimplementedTestplatformTestplanRelServer) PAGING(context.Context, *TestPlanCaseRelPagingRequest) (*TestPlanCaseRelPagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PAGING not implemented")
}
func (*UnimplementedTestplatformTestplanRelServer) REMOVE_ISSUE_RELATION(context.Context, *TestPlanCaseRelIssueRelationRemoveRequest) (*TestPlanCaseRelIssueRelationRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method REMOVE_ISSUE_RELATION not implemented")
}

func RegisterTestplatformTestplanRelServer(s grpc1.ServiceRegistrar, srv TestplatformTestplanRelServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TestplatformTestplanRel_serviceDesc(srv, opts...), srv)
}

var _TestplatformTestplanRel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.openapiv1.testplatform.testplatform_testplan_rel",
	HandlerType: (*TestplatformTestplanRelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "testplatform_testplan_rel.proto",
}

func _get_TestplatformTestplanRel_serviceDesc(srv TestplatformTestplanRelServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TestplatformTestplanRel_ADD_ISSUE_RELATION_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ADD_ISSUE_RELATION(ctx, req.(*TestPlanCaseRelIssueRelationAddRequest))
	}
	var _TestplatformTestplanRel_ADD_ISSUE_RELATION_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_ADD_ISSUE_RELATION_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "ADD_ISSUE_RELATION", srv)
		_TestplatformTestplanRel_ADD_ISSUE_RELATION_Handler = h.Interceptor(_TestplatformTestplanRel_ADD_ISSUE_RELATION_Handler)
	}

	_TestplatformTestplanRel_BATCH_UPDATE_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.BATCH_UPDATE(ctx, req.(*TestPlanCaseRelBatchUpdateRequest))
	}
	var _TestplatformTestplanRel_BATCH_UPDATE_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_BATCH_UPDATE_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "BATCH_UPDATE", srv)
		_TestplatformTestplanRel_BATCH_UPDATE_Handler = h.Interceptor(_TestplatformTestplanRel_BATCH_UPDATE_Handler)
	}

	_TestplatformTestplanRel_CREATE_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CREATE(ctx, req.(*TestPlanCaseRelCreateRequest))
	}
	var _TestplatformTestplanRel_CREATE_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_CREATE_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "CREATE", srv)
		_TestplatformTestplanRel_CREATE_Handler = h.Interceptor(_TestplatformTestplanRel_CREATE_Handler)
	}

	_TestplatformTestplanRel_EXPORT_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.EXPORT(ctx, req.(*TestPlanCaseRelExportRequest))
	}
	var _TestplatformTestplanRel_EXPORT_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_EXPORT_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "EXPORT", srv)
		_TestplatformTestplanRel_EXPORT_Handler = h.Interceptor(_TestplatformTestplanRel_EXPORT_Handler)
	}

	_TestplatformTestplanRel_GET_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GET(ctx, req.(*TestPlanCaseRelGetRequest))
	}
	var _TestplatformTestplanRel_GET_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_GET_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "GET", srv)
		_TestplatformTestplanRel_GET_Handler = h.Interceptor(_TestplatformTestplanRel_GET_Handler)
	}

	_TestplatformTestplanRel_PAGING_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.PAGING(ctx, req.(*TestPlanCaseRelPagingRequest))
	}
	var _TestplatformTestplanRel_PAGING_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_PAGING_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "PAGING", srv)
		_TestplatformTestplanRel_PAGING_Handler = h.Interceptor(_TestplatformTestplanRel_PAGING_Handler)
	}

	_TestplatformTestplanRel_REMOVE_ISSUE_RELATION_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.REMOVE_ISSUE_RELATION(ctx, req.(*TestPlanCaseRelIssueRelationRemoveRequest))
	}
	var _TestplatformTestplanRel_REMOVE_ISSUE_RELATION_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TestplatformTestplanRel_REMOVE_ISSUE_RELATION_info = transport.NewServiceInfo("erda.openapiv1.testplatform.testplatform_testplan_rel", "REMOVE_ISSUE_RELATION", srv)
		_TestplatformTestplanRel_REMOVE_ISSUE_RELATION_Handler = h.Interceptor(_TestplatformTestplanRel_REMOVE_ISSUE_RELATION_Handler)
	}

	var serviceDesc = _TestplatformTestplanRel_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ADD_ISSUE_RELATION",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelIssueRelationAddRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).ADD_ISSUE_RELATION(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_ADD_ISSUE_RELATION_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_ADD_ISSUE_RELATION_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/ADD_ISSUE_RELATION",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_ADD_ISSUE_RELATION_Handler)
			},
		},
		{
			MethodName: "BATCH_UPDATE",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelBatchUpdateRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).BATCH_UPDATE(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_BATCH_UPDATE_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_BATCH_UPDATE_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/BATCH_UPDATE",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_BATCH_UPDATE_Handler)
			},
		},
		{
			MethodName: "CREATE",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelCreateRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).CREATE(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_CREATE_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_CREATE_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/CREATE",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_CREATE_Handler)
			},
		},
		{
			MethodName: "EXPORT",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelExportRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).EXPORT(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_EXPORT_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_EXPORT_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/EXPORT",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_EXPORT_Handler)
			},
		},
		{
			MethodName: "GET",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelGetRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).GET(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_GET_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_GET_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/GET",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_GET_Handler)
			},
		},
		{
			MethodName: "PAGING",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelPagingRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).PAGING(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_PAGING_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_PAGING_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/PAGING",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_PAGING_Handler)
			},
		},
		{
			MethodName: "REMOVE_ISSUE_RELATION",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestPlanCaseRelIssueRelationRemoveRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TestplatformTestplanRelServer).REMOVE_ISSUE_RELATION(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TestplatformTestplanRel_REMOVE_ISSUE_RELATION_info)
				}
				if interceptor == nil {
					return _TestplatformTestplanRel_REMOVE_ISSUE_RELATION_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.testplatform.testplatform_testplan_rel/REMOVE_ISSUE_RELATION",
				}
				return interceptor(ctx, in, info, _TestplatformTestplanRel_REMOVE_ISSUE_RELATION_Handler)
			},
		},
	}
	return &serviceDesc
}
