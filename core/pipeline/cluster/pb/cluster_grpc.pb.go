// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: cluster.proto

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

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	ClusterHook(ctx context.Context, in *ClusterHookRequest, opts ...grpc.CallOption) (*ClusterHookResponse, error)
}

type clusterServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewClusterServiceClient(cc grpc1.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) ClusterHook(ctx context.Context, in *ClusterHookRequest, opts ...grpc.CallOption) (*ClusterHookResponse, error) {
	out := new(ClusterHookResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.cluster.ClusterService/ClusterHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	ClusterHook(context.Context, *ClusterHookRequest) (*ClusterHookResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (*UnimplementedClusterServiceServer) ClusterHook(context.Context, *ClusterHookRequest) (*ClusterHookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterHook not implemented")
}

func RegisterClusterServiceServer(s grpc1.ServiceRegistrar, srv ClusterServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ClusterService_serviceDesc(srv, opts...), srv)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.pipeline.cluster.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cluster.proto",
}

func _get_ClusterService_serviceDesc(srv ClusterServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ClusterService_ClusterHook_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ClusterHook(ctx, req.(*ClusterHookRequest))
	}
	var _ClusterService_ClusterHook_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_ClusterHook_info = transport.NewServiceInfo("erda.core.pipeline.cluster.ClusterService", "ClusterHook", srv)
		_ClusterService_ClusterHook_Handler = h.Interceptor(_ClusterService_ClusterHook_Handler)
	}

	var serviceDesc = _ClusterService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ClusterHook",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ClusterHookRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).ClusterHook(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_ClusterHook_info)
				}
				if interceptor == nil {
					return _ClusterService_ClusterHook_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.cluster.ClusterService/ClusterHook",
				}
				return interceptor(ctx, in, info, _ClusterService_ClusterHook_Handler)
			},
		},
	}
	return &serviceDesc
}
