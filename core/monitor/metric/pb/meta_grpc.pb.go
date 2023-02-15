// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: meta.proto

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

// MetricMetaServiceClient is the client API for MetricMetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricMetaServiceClient interface {
	ListMetricNames(ctx context.Context, in *ListMetricNamesRequest, opts ...grpc.CallOption) (*ListMetricNamesResponse, error)
	ListMetricMeta(ctx context.Context, in *ListMetricMetaRequest, opts ...grpc.CallOption) (*ListMetricMetaResponse, error)
	RegisterMetricMeta(ctx context.Context, in *RegisterMetricMetaRequest, opts ...grpc.CallOption) (*RegisterMetricMetaResponse, error)
	UnRegisterMetricMeta(ctx context.Context, in *UnRegisterMetricMetaRequest, opts ...grpc.CallOption) (*UnRegisterMetricMetaResponse, error)
	ListMetricGroups(ctx context.Context, in *ListMetricGroupsRequest, opts ...grpc.CallOption) (*ListMetricGroupsResponse, error)
	GetMetricGroup(ctx context.Context, in *GetMetricGroupRequest, opts ...grpc.CallOption) (*GetMetricGroupResponse, error)
}

type metricMetaServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewMetricMetaServiceClient(cc grpc1.ClientConnInterface) MetricMetaServiceClient {
	return &metricMetaServiceClient{cc}
}

func (c *metricMetaServiceClient) ListMetricNames(ctx context.Context, in *ListMetricNamesRequest, opts ...grpc.CallOption) (*ListMetricNamesResponse, error) {
	out := new(ListMetricNamesResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.metric.MetricMetaService/ListMetricNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricMetaServiceClient) ListMetricMeta(ctx context.Context, in *ListMetricMetaRequest, opts ...grpc.CallOption) (*ListMetricMetaResponse, error) {
	out := new(ListMetricMetaResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.metric.MetricMetaService/ListMetricMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricMetaServiceClient) RegisterMetricMeta(ctx context.Context, in *RegisterMetricMetaRequest, opts ...grpc.CallOption) (*RegisterMetricMetaResponse, error) {
	out := new(RegisterMetricMetaResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.metric.MetricMetaService/RegisterMetricMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricMetaServiceClient) UnRegisterMetricMeta(ctx context.Context, in *UnRegisterMetricMetaRequest, opts ...grpc.CallOption) (*UnRegisterMetricMetaResponse, error) {
	out := new(UnRegisterMetricMetaResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.metric.MetricMetaService/UnRegisterMetricMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricMetaServiceClient) ListMetricGroups(ctx context.Context, in *ListMetricGroupsRequest, opts ...grpc.CallOption) (*ListMetricGroupsResponse, error) {
	out := new(ListMetricGroupsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.metric.MetricMetaService/ListMetricGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricMetaServiceClient) GetMetricGroup(ctx context.Context, in *GetMetricGroupRequest, opts ...grpc.CallOption) (*GetMetricGroupResponse, error) {
	out := new(GetMetricGroupResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.metric.MetricMetaService/GetMetricGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricMetaServiceServer is the server API for MetricMetaService service.
// All implementations should embed UnimplementedMetricMetaServiceServer
// for forward compatibility
type MetricMetaServiceServer interface {
	ListMetricNames(context.Context, *ListMetricNamesRequest) (*ListMetricNamesResponse, error)
	ListMetricMeta(context.Context, *ListMetricMetaRequest) (*ListMetricMetaResponse, error)
	RegisterMetricMeta(context.Context, *RegisterMetricMetaRequest) (*RegisterMetricMetaResponse, error)
	UnRegisterMetricMeta(context.Context, *UnRegisterMetricMetaRequest) (*UnRegisterMetricMetaResponse, error)
	ListMetricGroups(context.Context, *ListMetricGroupsRequest) (*ListMetricGroupsResponse, error)
	GetMetricGroup(context.Context, *GetMetricGroupRequest) (*GetMetricGroupResponse, error)
}

// UnimplementedMetricMetaServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMetricMetaServiceServer struct {
}

func (*UnimplementedMetricMetaServiceServer) ListMetricNames(context.Context, *ListMetricNamesRequest) (*ListMetricNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricNames not implemented")
}
func (*UnimplementedMetricMetaServiceServer) ListMetricMeta(context.Context, *ListMetricMetaRequest) (*ListMetricMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricMeta not implemented")
}
func (*UnimplementedMetricMetaServiceServer) RegisterMetricMeta(context.Context, *RegisterMetricMetaRequest) (*RegisterMetricMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMetricMeta not implemented")
}
func (*UnimplementedMetricMetaServiceServer) UnRegisterMetricMeta(context.Context, *UnRegisterMetricMetaRequest) (*UnRegisterMetricMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterMetricMeta not implemented")
}
func (*UnimplementedMetricMetaServiceServer) ListMetricGroups(context.Context, *ListMetricGroupsRequest) (*ListMetricGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricGroups not implemented")
}
func (*UnimplementedMetricMetaServiceServer) GetMetricGroup(context.Context, *GetMetricGroupRequest) (*GetMetricGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricGroup not implemented")
}

func RegisterMetricMetaServiceServer(s grpc1.ServiceRegistrar, srv MetricMetaServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_MetricMetaService_serviceDesc(srv, opts...), srv)
}

var _MetricMetaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.monitor.metric.MetricMetaService",
	HandlerType: (*MetricMetaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "meta.proto",
}

func _get_MetricMetaService_serviceDesc(srv MetricMetaServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_MetricMetaService_ListMetricNames_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListMetricNames(ctx, req.(*ListMetricNamesRequest))
	}
	var _MetricMetaService_ListMetricNames_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MetricMetaService_ListMetricNames_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricMetaService", "ListMetricNames", srv)
		_MetricMetaService_ListMetricNames_Handler = h.Interceptor(_MetricMetaService_ListMetricNames_Handler)
	}

	_MetricMetaService_ListMetricMeta_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListMetricMeta(ctx, req.(*ListMetricMetaRequest))
	}
	var _MetricMetaService_ListMetricMeta_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MetricMetaService_ListMetricMeta_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricMetaService", "ListMetricMeta", srv)
		_MetricMetaService_ListMetricMeta_Handler = h.Interceptor(_MetricMetaService_ListMetricMeta_Handler)
	}

	_MetricMetaService_RegisterMetricMeta_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RegisterMetricMeta(ctx, req.(*RegisterMetricMetaRequest))
	}
	var _MetricMetaService_RegisterMetricMeta_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MetricMetaService_RegisterMetricMeta_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricMetaService", "RegisterMetricMeta", srv)
		_MetricMetaService_RegisterMetricMeta_Handler = h.Interceptor(_MetricMetaService_RegisterMetricMeta_Handler)
	}

	_MetricMetaService_UnRegisterMetricMeta_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UnRegisterMetricMeta(ctx, req.(*UnRegisterMetricMetaRequest))
	}
	var _MetricMetaService_UnRegisterMetricMeta_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MetricMetaService_UnRegisterMetricMeta_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricMetaService", "UnRegisterMetricMeta", srv)
		_MetricMetaService_UnRegisterMetricMeta_Handler = h.Interceptor(_MetricMetaService_UnRegisterMetricMeta_Handler)
	}

	_MetricMetaService_ListMetricGroups_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListMetricGroups(ctx, req.(*ListMetricGroupsRequest))
	}
	var _MetricMetaService_ListMetricGroups_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MetricMetaService_ListMetricGroups_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricMetaService", "ListMetricGroups", srv)
		_MetricMetaService_ListMetricGroups_Handler = h.Interceptor(_MetricMetaService_ListMetricGroups_Handler)
	}

	_MetricMetaService_GetMetricGroup_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetMetricGroup(ctx, req.(*GetMetricGroupRequest))
	}
	var _MetricMetaService_GetMetricGroup_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MetricMetaService_GetMetricGroup_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricMetaService", "GetMetricGroup", srv)
		_MetricMetaService_GetMetricGroup_Handler = h.Interceptor(_MetricMetaService_GetMetricGroup_Handler)
	}

	var serviceDesc = _MetricMetaService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ListMetricNames",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListMetricNamesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MetricMetaServiceServer).ListMetricNames(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MetricMetaService_ListMetricNames_info)
				}
				if interceptor == nil {
					return _MetricMetaService_ListMetricNames_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.metric.MetricMetaService/ListMetricNames",
				}
				return interceptor(ctx, in, info, _MetricMetaService_ListMetricNames_Handler)
			},
		},
		{
			MethodName: "ListMetricMeta",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListMetricMetaRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MetricMetaServiceServer).ListMetricMeta(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MetricMetaService_ListMetricMeta_info)
				}
				if interceptor == nil {
					return _MetricMetaService_ListMetricMeta_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.metric.MetricMetaService/ListMetricMeta",
				}
				return interceptor(ctx, in, info, _MetricMetaService_ListMetricMeta_Handler)
			},
		},
		{
			MethodName: "RegisterMetricMeta",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RegisterMetricMetaRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MetricMetaServiceServer).RegisterMetricMeta(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MetricMetaService_RegisterMetricMeta_info)
				}
				if interceptor == nil {
					return _MetricMetaService_RegisterMetricMeta_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.metric.MetricMetaService/RegisterMetricMeta",
				}
				return interceptor(ctx, in, info, _MetricMetaService_RegisterMetricMeta_Handler)
			},
		},
		{
			MethodName: "UnRegisterMetricMeta",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UnRegisterMetricMetaRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MetricMetaServiceServer).UnRegisterMetricMeta(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MetricMetaService_UnRegisterMetricMeta_info)
				}
				if interceptor == nil {
					return _MetricMetaService_UnRegisterMetricMeta_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.metric.MetricMetaService/UnRegisterMetricMeta",
				}
				return interceptor(ctx, in, info, _MetricMetaService_UnRegisterMetricMeta_Handler)
			},
		},
		{
			MethodName: "ListMetricGroups",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListMetricGroupsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MetricMetaServiceServer).ListMetricGroups(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MetricMetaService_ListMetricGroups_info)
				}
				if interceptor == nil {
					return _MetricMetaService_ListMetricGroups_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.metric.MetricMetaService/ListMetricGroups",
				}
				return interceptor(ctx, in, info, _MetricMetaService_ListMetricGroups_Handler)
			},
		},
		{
			MethodName: "GetMetricGroup",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetMetricGroupRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MetricMetaServiceServer).GetMetricGroup(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MetricMetaService_GetMetricGroup_info)
				}
				if interceptor == nil {
					return _MetricMetaService_GetMetricGroup_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.metric.MetricMetaService/GetMetricGroup",
				}
				return interceptor(ctx, in, info, _MetricMetaService_GetMetricGroup_Handler)
			},
		},
	}
	return &serviceDesc
}
