// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: msp_apm_block.proto

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

// MspApmBlockClient is the client API for MspApmBlock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MspApmBlockClient interface {
	CREATE_BLOCK(ctx context.Context, in *CREATE_BLOCK_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DELETE_BLOCK(ctx context.Context, in *DELETE_BLOCK_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GET_BLOCK(ctx context.Context, in *GET_BLOCK_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LIST_BLOCKS(ctx context.Context, in *LIST_BLOCKS_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TMC_METRIC_DASHBOARD_UPDATE(ctx context.Context, in *TMC_METRIC_DASHBOARD_UPDATE_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mspApmBlockClient struct {
	cc grpc1.ClientConnInterface
}

func NewMspApmBlockClient(cc grpc1.ClientConnInterface) MspApmBlockClient {
	return &mspApmBlockClient{cc}
}

func (c *mspApmBlockClient) CREATE_BLOCK(ctx context.Context, in *CREATE_BLOCK_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_apm_block/CREATE_BLOCK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mspApmBlockClient) DELETE_BLOCK(ctx context.Context, in *DELETE_BLOCK_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_apm_block/DELETE_BLOCK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mspApmBlockClient) GET_BLOCK(ctx context.Context, in *GET_BLOCK_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_apm_block/GET_BLOCK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mspApmBlockClient) LIST_BLOCKS(ctx context.Context, in *LIST_BLOCKS_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_apm_block/LIST_BLOCKS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mspApmBlockClient) TMC_METRIC_DASHBOARD_UPDATE(ctx context.Context, in *TMC_METRIC_DASHBOARD_UPDATE_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_apm_block/TMC_METRIC_DASHBOARD_UPDATE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MspApmBlockServer is the server API for MspApmBlock service.
// All implementations should embed UnimplementedMspApmBlockServer
// for forward compatibility
type MspApmBlockServer interface {
	CREATE_BLOCK(context.Context, *CREATE_BLOCK_Request) (*emptypb.Empty, error)
	DELETE_BLOCK(context.Context, *DELETE_BLOCK_Request) (*emptypb.Empty, error)
	GET_BLOCK(context.Context, *GET_BLOCK_Request) (*emptypb.Empty, error)
	LIST_BLOCKS(context.Context, *LIST_BLOCKS_Request) (*emptypb.Empty, error)
	TMC_METRIC_DASHBOARD_UPDATE(context.Context, *TMC_METRIC_DASHBOARD_UPDATE_Request) (*emptypb.Empty, error)
}

// UnimplementedMspApmBlockServer should be embedded to have forward compatible implementations.
type UnimplementedMspApmBlockServer struct {
}

func (*UnimplementedMspApmBlockServer) CREATE_BLOCK(context.Context, *CREATE_BLOCK_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CREATE_BLOCK not implemented")
}
func (*UnimplementedMspApmBlockServer) DELETE_BLOCK(context.Context, *DELETE_BLOCK_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DELETE_BLOCK not implemented")
}
func (*UnimplementedMspApmBlockServer) GET_BLOCK(context.Context, *GET_BLOCK_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GET_BLOCK not implemented")
}
func (*UnimplementedMspApmBlockServer) LIST_BLOCKS(context.Context, *LIST_BLOCKS_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LIST_BLOCKS not implemented")
}
func (*UnimplementedMspApmBlockServer) TMC_METRIC_DASHBOARD_UPDATE(context.Context, *TMC_METRIC_DASHBOARD_UPDATE_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TMC_METRIC_DASHBOARD_UPDATE not implemented")
}

func RegisterMspApmBlockServer(s grpc1.ServiceRegistrar, srv MspApmBlockServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_MspApmBlock_serviceDesc(srv, opts...), srv)
}

var _MspApmBlock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.openapiv1.msp.msp_apm_block",
	HandlerType: (*MspApmBlockServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "msp_apm_block.proto",
}

func _get_MspApmBlock_serviceDesc(srv MspApmBlockServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_MspApmBlock_CREATE_BLOCK_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CREATE_BLOCK(ctx, req.(*CREATE_BLOCK_Request))
	}
	var _MspApmBlock_CREATE_BLOCK_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspApmBlock_CREATE_BLOCK_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "CREATE_BLOCK", srv)
		_MspApmBlock_CREATE_BLOCK_Handler = h.Interceptor(_MspApmBlock_CREATE_BLOCK_Handler)
	}

	_MspApmBlock_DELETE_BLOCK_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DELETE_BLOCK(ctx, req.(*DELETE_BLOCK_Request))
	}
	var _MspApmBlock_DELETE_BLOCK_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspApmBlock_DELETE_BLOCK_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "DELETE_BLOCK", srv)
		_MspApmBlock_DELETE_BLOCK_Handler = h.Interceptor(_MspApmBlock_DELETE_BLOCK_Handler)
	}

	_MspApmBlock_GET_BLOCK_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GET_BLOCK(ctx, req.(*GET_BLOCK_Request))
	}
	var _MspApmBlock_GET_BLOCK_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspApmBlock_GET_BLOCK_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "GET_BLOCK", srv)
		_MspApmBlock_GET_BLOCK_Handler = h.Interceptor(_MspApmBlock_GET_BLOCK_Handler)
	}

	_MspApmBlock_LIST_BLOCKS_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.LIST_BLOCKS(ctx, req.(*LIST_BLOCKS_Request))
	}
	var _MspApmBlock_LIST_BLOCKS_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspApmBlock_LIST_BLOCKS_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "LIST_BLOCKS", srv)
		_MspApmBlock_LIST_BLOCKS_Handler = h.Interceptor(_MspApmBlock_LIST_BLOCKS_Handler)
	}

	_MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.TMC_METRIC_DASHBOARD_UPDATE(ctx, req.(*TMC_METRIC_DASHBOARD_UPDATE_Request))
	}
	var _MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_apm_block", "TMC_METRIC_DASHBOARD_UPDATE", srv)
		_MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_Handler = h.Interceptor(_MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_Handler)
	}

	var serviceDesc = _MspApmBlock_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CREATE_BLOCK",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CREATE_BLOCK_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspApmBlockServer).CREATE_BLOCK(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspApmBlock_CREATE_BLOCK_info)
				}
				if interceptor == nil {
					return _MspApmBlock_CREATE_BLOCK_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_apm_block/CREATE_BLOCK",
				}
				return interceptor(ctx, in, info, _MspApmBlock_CREATE_BLOCK_Handler)
			},
		},
		{
			MethodName: "DELETE_BLOCK",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DELETE_BLOCK_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspApmBlockServer).DELETE_BLOCK(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspApmBlock_DELETE_BLOCK_info)
				}
				if interceptor == nil {
					return _MspApmBlock_DELETE_BLOCK_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_apm_block/DELETE_BLOCK",
				}
				return interceptor(ctx, in, info, _MspApmBlock_DELETE_BLOCK_Handler)
			},
		},
		{
			MethodName: "GET_BLOCK",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GET_BLOCK_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspApmBlockServer).GET_BLOCK(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspApmBlock_GET_BLOCK_info)
				}
				if interceptor == nil {
					return _MspApmBlock_GET_BLOCK_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_apm_block/GET_BLOCK",
				}
				return interceptor(ctx, in, info, _MspApmBlock_GET_BLOCK_Handler)
			},
		},
		{
			MethodName: "LIST_BLOCKS",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(LIST_BLOCKS_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspApmBlockServer).LIST_BLOCKS(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspApmBlock_LIST_BLOCKS_info)
				}
				if interceptor == nil {
					return _MspApmBlock_LIST_BLOCKS_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_apm_block/LIST_BLOCKS",
				}
				return interceptor(ctx, in, info, _MspApmBlock_LIST_BLOCKS_Handler)
			},
		},
		{
			MethodName: "TMC_METRIC_DASHBOARD_UPDATE",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TMC_METRIC_DASHBOARD_UPDATE_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspApmBlockServer).TMC_METRIC_DASHBOARD_UPDATE(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_info)
				}
				if interceptor == nil {
					return _MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_apm_block/TMC_METRIC_DASHBOARD_UPDATE",
				}
				return interceptor(ctx, in, info, _MspApmBlock_TMC_METRIC_DASHBOARD_UPDATE_Handler)
			},
		},
	}
	return &serviceDesc
}
