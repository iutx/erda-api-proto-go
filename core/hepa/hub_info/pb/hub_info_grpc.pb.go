// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: hub_info.proto

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

// HubInfoServiceClient is the client API for HubInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubInfoServiceClient interface {
	CreateOrGetHubInfo(ctx context.Context, in *CreateHubInfoReq, opts ...grpc.CallOption) (*GetHubInfoResp, error)
	GetHubInfo(ctx context.Context, in *GetHubInfoReq, opts ...grpc.CallOption) (*GetHubInfoResp, error)
}

type hubInfoServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewHubInfoServiceClient(cc grpc1.ClientConnInterface) HubInfoServiceClient {
	return &hubInfoServiceClient{cc}
}

func (c *hubInfoServiceClient) CreateOrGetHubInfo(ctx context.Context, in *CreateHubInfoReq, opts ...grpc.CallOption) (*GetHubInfoResp, error) {
	out := new(GetHubInfoResp)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.hub_info.HubInfoService/CreateOrGetHubInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubInfoServiceClient) GetHubInfo(ctx context.Context, in *GetHubInfoReq, opts ...grpc.CallOption) (*GetHubInfoResp, error) {
	out := new(GetHubInfoResp)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.hub_info.HubInfoService/GetHubInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubInfoServiceServer is the server API for HubInfoService service.
// All implementations should embed UnimplementedHubInfoServiceServer
// for forward compatibility
type HubInfoServiceServer interface {
	CreateOrGetHubInfo(context.Context, *CreateHubInfoReq) (*GetHubInfoResp, error)
	GetHubInfo(context.Context, *GetHubInfoReq) (*GetHubInfoResp, error)
}

// UnimplementedHubInfoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHubInfoServiceServer struct {
}

func (*UnimplementedHubInfoServiceServer) CreateOrGetHubInfo(context.Context, *CreateHubInfoReq) (*GetHubInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrGetHubInfo not implemented")
}
func (*UnimplementedHubInfoServiceServer) GetHubInfo(context.Context, *GetHubInfoReq) (*GetHubInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHubInfo not implemented")
}

func RegisterHubInfoServiceServer(s grpc1.ServiceRegistrar, srv HubInfoServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_HubInfoService_serviceDesc(srv, opts...), srv)
}

var _HubInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.hepa.hub_info.HubInfoService",
	HandlerType: (*HubInfoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "hub_info.proto",
}

func _get_HubInfoService_serviceDesc(srv HubInfoServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_HubInfoService_CreateOrGetHubInfo_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateOrGetHubInfo(ctx, req.(*CreateHubInfoReq))
	}
	var _HubInfoService_CreateOrGetHubInfo_info transport.ServiceInfo
	if h.Interceptor != nil {
		_HubInfoService_CreateOrGetHubInfo_info = transport.NewServiceInfo("erda.core.hepa.hub_info.HubInfoService", "CreateOrGetHubInfo", srv)
		_HubInfoService_CreateOrGetHubInfo_Handler = h.Interceptor(_HubInfoService_CreateOrGetHubInfo_Handler)
	}

	_HubInfoService_GetHubInfo_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetHubInfo(ctx, req.(*GetHubInfoReq))
	}
	var _HubInfoService_GetHubInfo_info transport.ServiceInfo
	if h.Interceptor != nil {
		_HubInfoService_GetHubInfo_info = transport.NewServiceInfo("erda.core.hepa.hub_info.HubInfoService", "GetHubInfo", srv)
		_HubInfoService_GetHubInfo_Handler = h.Interceptor(_HubInfoService_GetHubInfo_Handler)
	}

	var serviceDesc = _HubInfoService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CreateOrGetHubInfo",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateHubInfoReq)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(HubInfoServiceServer).CreateOrGetHubInfo(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _HubInfoService_CreateOrGetHubInfo_info)
				}
				if interceptor == nil {
					return _HubInfoService_CreateOrGetHubInfo_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.hub_info.HubInfoService/CreateOrGetHubInfo",
				}
				return interceptor(ctx, in, info, _HubInfoService_CreateOrGetHubInfo_Handler)
			},
		},
		{
			MethodName: "GetHubInfo",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetHubInfoReq)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(HubInfoServiceServer).GetHubInfo(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _HubInfoService_GetHubInfo_info)
				}
				if interceptor == nil {
					return _HubInfoService_GetHubInfo_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.hub_info.HubInfoService/GetHubInfo",
				}
				return interceptor(ctx, in, info, _HubInfoService_GetHubInfo_Handler)
			},
		},
	}
	return &serviceDesc
}
