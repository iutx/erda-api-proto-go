// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: msp_tenant.proto

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

// MspTenantClient is the client API for MspTenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MspTenantClient interface {
	CREATE_TENANTS(ctx context.Context, in *CREATE_TENANTS_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GET_TENANTS(ctx context.Context, in *GET_TENANTS_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mspTenantClient struct {
	cc grpc1.ClientConnInterface
}

func NewMspTenantClient(cc grpc1.ClientConnInterface) MspTenantClient {
	return &mspTenantClient{cc}
}

func (c *mspTenantClient) CREATE_TENANTS(ctx context.Context, in *CREATE_TENANTS_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_tenant/CREATE_TENANTS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mspTenantClient) GET_TENANTS(ctx context.Context, in *GET_TENANTS_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.msp.msp_tenant/GET_TENANTS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MspTenantServer is the server API for MspTenant service.
// All implementations should embed UnimplementedMspTenantServer
// for forward compatibility
type MspTenantServer interface {
	CREATE_TENANTS(context.Context, *CREATE_TENANTS_Request) (*emptypb.Empty, error)
	GET_TENANTS(context.Context, *GET_TENANTS_Request) (*emptypb.Empty, error)
}

// UnimplementedMspTenantServer should be embedded to have forward compatible implementations.
type UnimplementedMspTenantServer struct {
}

func (*UnimplementedMspTenantServer) CREATE_TENANTS(context.Context, *CREATE_TENANTS_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CREATE_TENANTS not implemented")
}
func (*UnimplementedMspTenantServer) GET_TENANTS(context.Context, *GET_TENANTS_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GET_TENANTS not implemented")
}

func RegisterMspTenantServer(s grpc1.ServiceRegistrar, srv MspTenantServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_MspTenant_serviceDesc(srv, opts...), srv)
}

var _MspTenant_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.openapiv1.msp.msp_tenant",
	HandlerType: (*MspTenantServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "msp_tenant.proto",
}

func _get_MspTenant_serviceDesc(srv MspTenantServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_MspTenant_CREATE_TENANTS_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CREATE_TENANTS(ctx, req.(*CREATE_TENANTS_Request))
	}
	var _MspTenant_CREATE_TENANTS_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspTenant_CREATE_TENANTS_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_tenant", "CREATE_TENANTS", srv)
		_MspTenant_CREATE_TENANTS_Handler = h.Interceptor(_MspTenant_CREATE_TENANTS_Handler)
	}

	_MspTenant_GET_TENANTS_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GET_TENANTS(ctx, req.(*GET_TENANTS_Request))
	}
	var _MspTenant_GET_TENANTS_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MspTenant_GET_TENANTS_info = transport.NewServiceInfo("erda.openapiv1.msp.msp_tenant", "GET_TENANTS", srv)
		_MspTenant_GET_TENANTS_Handler = h.Interceptor(_MspTenant_GET_TENANTS_Handler)
	}

	var serviceDesc = _MspTenant_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CREATE_TENANTS",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CREATE_TENANTS_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspTenantServer).CREATE_TENANTS(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspTenant_CREATE_TENANTS_info)
				}
				if interceptor == nil {
					return _MspTenant_CREATE_TENANTS_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_tenant/CREATE_TENANTS",
				}
				return interceptor(ctx, in, info, _MspTenant_CREATE_TENANTS_Handler)
			},
		},
		{
			MethodName: "GET_TENANTS",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GET_TENANTS_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MspTenantServer).GET_TENANTS(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MspTenant_GET_TENANTS_info)
				}
				if interceptor == nil {
					return _MspTenant_GET_TENANTS_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.msp.msp_tenant/GET_TENANTS",
				}
				return interceptor(ctx, in, info, _MspTenant_GET_TENANTS_Handler)
			},
		},
	}
	return &serviceDesc
}
