// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: tenant.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantServiceClient interface {
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error)
	GetTenantProject(ctx context.Context, in *GetTenantProjectRequest, opts ...grpc.CallOption) (*GetTenantProjectResponse, error)
}

type tenantServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewTenantServiceClient(cc grpc1.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.tenant.TenantService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.tenant.TenantService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error) {
	out := new(DeleteTenantResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.tenant.TenantService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenantProject(ctx context.Context, in *GetTenantProjectRequest, opts ...grpc.CallOption) (*GetTenantProjectResponse, error) {
	out := new(GetTenantProjectResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.tenant.TenantService/GetTenantProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations should embed UnimplementedTenantServiceServer
// for forward compatibility
type TenantServiceServer interface {
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error)
	DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error)
	GetTenantProject(context.Context, *GetTenantProjectRequest) (*GetTenantProjectResponse, error)
}

// UnimplementedTenantServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (*UnimplementedTenantServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (*UnimplementedTenantServiceServer) GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (*UnimplementedTenantServiceServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (*UnimplementedTenantServiceServer) GetTenantProject(context.Context, *GetTenantProjectRequest) (*GetTenantProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantProject not implemented")
}

func RegisterTenantServiceServer(s grpc1.ServiceRegistrar, srv TenantServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TenantService_serviceDesc(srv, opts...), srv)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.msp.tenant.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "tenant.proto",
}

func _get_TenantService_serviceDesc(srv TenantServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TenantService_CreateTenant_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	var _TenantService_CreateTenant_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TenantService_CreateTenant_info = transport.NewServiceInfo("erda.msp.tenant.TenantService", "CreateTenant", srv)
		_TenantService_CreateTenant_Handler = h.Interceptor(_TenantService_CreateTenant_Handler)
	}

	_TenantService_GetTenant_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTenant(ctx, req.(*GetTenantRequest))
	}
	var _TenantService_GetTenant_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TenantService_GetTenant_info = transport.NewServiceInfo("erda.msp.tenant.TenantService", "GetTenant", srv)
		_TenantService_GetTenant_Handler = h.Interceptor(_TenantService_GetTenant_Handler)
	}

	_TenantService_DeleteTenant_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	var _TenantService_DeleteTenant_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TenantService_DeleteTenant_info = transport.NewServiceInfo("erda.msp.tenant.TenantService", "DeleteTenant", srv)
		_TenantService_DeleteTenant_Handler = h.Interceptor(_TenantService_DeleteTenant_Handler)
	}

	_TenantService_GetTenantProject_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTenantProject(ctx, req.(*GetTenantProjectRequest))
	}
	var _TenantService_GetTenantProject_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TenantService_GetTenantProject_info = transport.NewServiceInfo("erda.msp.tenant.TenantService", "GetTenantProject", srv)
		_TenantService_GetTenantProject_Handler = h.Interceptor(_TenantService_GetTenantProject_Handler)
	}

	var serviceDesc = _TenantService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateTenantRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TenantServiceServer).CreateTenant(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TenantService_CreateTenant_info)
				}
				if interceptor == nil {
					return _TenantService_CreateTenant_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.tenant.TenantService/CreateTenant",
				}
				return interceptor(ctx, in, info, _TenantService_CreateTenant_Handler)
			},
		},
		{
			MethodName: "GetTenant",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTenantRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TenantServiceServer).GetTenant(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TenantService_GetTenant_info)
				}
				if interceptor == nil {
					return _TenantService_GetTenant_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.tenant.TenantService/GetTenant",
				}
				return interceptor(ctx, in, info, _TenantService_GetTenant_Handler)
			},
		},
		{
			MethodName: "DeleteTenant",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteTenantRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TenantServiceServer).DeleteTenant(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TenantService_DeleteTenant_info)
				}
				if interceptor == nil {
					return _TenantService_DeleteTenant_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.tenant.TenantService/DeleteTenant",
				}
				return interceptor(ctx, in, info, _TenantService_DeleteTenant_Handler)
			},
		},
		{
			MethodName: "GetTenantProject",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTenantProjectRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TenantServiceServer).GetTenantProject(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TenantService_GetTenantProject_info)
				}
				if interceptor == nil {
					return _TenantService_GetTenantProject_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.tenant.TenantService/GetTenantProject",
				}
				return interceptor(ctx, in, info, _TenantService_GetTenantProject_Handler)
			},
		},
	}
	return &serviceDesc
}
