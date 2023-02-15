// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: cp.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// OpenapiComponentProtocolClient is the client API for OpenapiComponentProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenapiComponentProtocolClient interface {
	Render(ctx context.Context, in *structpb.Value, opts ...grpc.CallOption) (*structpb.Value, error)
}

type openapiComponentProtocolClient struct {
	cc grpc1.ClientConnInterface
}

func NewOpenapiComponentProtocolClient(cc grpc1.ClientConnInterface) OpenapiComponentProtocolClient {
	return &openapiComponentProtocolClient{cc}
}

func (c *openapiComponentProtocolClient) Render(ctx context.Context, in *structpb.Value, opts ...grpc.CallOption) (*structpb.Value, error) {
	out := new(structpb.Value)
	err := c.cc.Invoke(ctx, "/erda.core.openapi.component_protocol.openapi_component_protocol/Render", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenapiComponentProtocolServer is the server API for OpenapiComponentProtocol service.
// All implementations should embed UnimplementedOpenapiComponentProtocolServer
// for forward compatibility
type OpenapiComponentProtocolServer interface {
	Render(context.Context, *structpb.Value) (*structpb.Value, error)
}

// UnimplementedOpenapiComponentProtocolServer should be embedded to have forward compatible implementations.
type UnimplementedOpenapiComponentProtocolServer struct {
}

func (*UnimplementedOpenapiComponentProtocolServer) Render(context.Context, *structpb.Value) (*structpb.Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Render not implemented")
}

func RegisterOpenapiComponentProtocolServer(s grpc1.ServiceRegistrar, srv OpenapiComponentProtocolServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_OpenapiComponentProtocol_serviceDesc(srv, opts...), srv)
}

var _OpenapiComponentProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.openapi.component_protocol.openapi_component_protocol",
	HandlerType: (*OpenapiComponentProtocolServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cp.proto",
}

func _get_OpenapiComponentProtocol_serviceDesc(srv OpenapiComponentProtocolServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_OpenapiComponentProtocol_Render_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Render(ctx, req.(*structpb.Value))
	}
	var _OpenapiComponentProtocol_Render_info transport.ServiceInfo
	if h.Interceptor != nil {
		_OpenapiComponentProtocol_Render_info = transport.NewServiceInfo("erda.core.openapi.component_protocol.openapi_component_protocol", "Render", srv)
		_OpenapiComponentProtocol_Render_Handler = h.Interceptor(_OpenapiComponentProtocol_Render_Handler)
	}

	var serviceDesc = _OpenapiComponentProtocol_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Render",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(structpb.Value)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(OpenapiComponentProtocolServer).Render(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _OpenapiComponentProtocol_Render_info)
				}
				if interceptor == nil {
					return _OpenapiComponentProtocol_Render_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.openapi.component_protocol.openapi_component_protocol/Render",
				}
				return interceptor(ctx, in, info, _OpenapiComponentProtocol_Render_Handler)
			},
		},
	}
	return &serviceDesc
}