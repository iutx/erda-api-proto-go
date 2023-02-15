// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: dop_config_manage.proto

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

// DopConfigManageClient is the client API for DopConfigManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DopConfigManageClient interface {
	CONFIG_MANAGE_CONFIG_DEL(ctx context.Context, in *CONFIG_MANAGE_CONFIG_DEL_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_EXPORT(ctx context.Context, in *CONFIG_MANAGE_CONFIG_EXPORT_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_GET(ctx context.Context, in *CONFIG_MANAGE_CONFIG_GET_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_IMPORT(ctx context.Context, in *CONFIG_MANAGE_CONFIG_IMPORT_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_POST(ctx context.Context, in *CONFIG_MANAGE_CONFIG_POST_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_PUT(ctx context.Context, in *CONFIG_MANAGE_CONFIG_PUT_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_DEPLOY_CONFIG_GET(ctx context.Context, in *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_MULTI_NS_CONFIG_GET(ctx context.Context, in *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_NAMESPACE_DEL(ctx context.Context, in *CONFIG_MANAGE_NAMESPACE_DEL_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_NAMESPACE_FIX(ctx context.Context, in *CONFIG_MANAGE_NAMESPACE_FIX_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CONFIG_MANAGE_NAMESPACE_POST(ctx context.Context, in *CONFIG_MANAGE_NAMESPACE_POST_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dopConfigManageClient struct {
	cc grpc1.ClientConnInterface
}

func NewDopConfigManageClient(cc grpc1.ClientConnInterface) DopConfigManageClient {
	return &dopConfigManageClient{cc}
}

func (c *dopConfigManageClient) CONFIG_MANAGE_CONFIG_DEL(ctx context.Context, in *CONFIG_MANAGE_CONFIG_DEL_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_DEL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_CONFIG_EXPORT(ctx context.Context, in *CONFIG_MANAGE_CONFIG_EXPORT_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_EXPORT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_CONFIG_GET(ctx context.Context, in *CONFIG_MANAGE_CONFIG_GET_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_GET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_CONFIG_IMPORT(ctx context.Context, in *CONFIG_MANAGE_CONFIG_IMPORT_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_IMPORT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_CONFIG_POST(ctx context.Context, in *CONFIG_MANAGE_CONFIG_POST_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_POST", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_CONFIG_PUT(ctx context.Context, in *CONFIG_MANAGE_CONFIG_PUT_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_PUT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_DEPLOY_CONFIG_GET(ctx context.Context, in *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_DEPLOY_CONFIG_GET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_MULTI_NS_CONFIG_GET(ctx context.Context, in *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_MULTI_NS_CONFIG_GET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_NAMESPACE_DEL(ctx context.Context, in *CONFIG_MANAGE_NAMESPACE_DEL_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_NAMESPACE_DEL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_NAMESPACE_FIX(ctx context.Context, in *CONFIG_MANAGE_NAMESPACE_FIX_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_NAMESPACE_FIX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopConfigManageClient) CONFIG_MANAGE_NAMESPACE_POST(ctx context.Context, in *CONFIG_MANAGE_NAMESPACE_POST_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_NAMESPACE_POST", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DopConfigManageServer is the server API for DopConfigManage service.
// All implementations should embed UnimplementedDopConfigManageServer
// for forward compatibility
type DopConfigManageServer interface {
	CONFIG_MANAGE_CONFIG_DEL(context.Context, *CONFIG_MANAGE_CONFIG_DEL_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_EXPORT(context.Context, *CONFIG_MANAGE_CONFIG_EXPORT_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_GET(context.Context, *CONFIG_MANAGE_CONFIG_GET_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_IMPORT(context.Context, *CONFIG_MANAGE_CONFIG_IMPORT_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_POST(context.Context, *CONFIG_MANAGE_CONFIG_POST_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_CONFIG_PUT(context.Context, *CONFIG_MANAGE_CONFIG_PUT_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_DEPLOY_CONFIG_GET(context.Context, *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_MULTI_NS_CONFIG_GET(context.Context, *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_NAMESPACE_DEL(context.Context, *CONFIG_MANAGE_NAMESPACE_DEL_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_NAMESPACE_FIX(context.Context, *CONFIG_MANAGE_NAMESPACE_FIX_Request) (*emptypb.Empty, error)
	CONFIG_MANAGE_NAMESPACE_POST(context.Context, *CONFIG_MANAGE_NAMESPACE_POST_Request) (*emptypb.Empty, error)
}

// UnimplementedDopConfigManageServer should be embedded to have forward compatible implementations.
type UnimplementedDopConfigManageServer struct {
}

func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_CONFIG_DEL(context.Context, *CONFIG_MANAGE_CONFIG_DEL_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_CONFIG_DEL not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_CONFIG_EXPORT(context.Context, *CONFIG_MANAGE_CONFIG_EXPORT_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_CONFIG_EXPORT not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_CONFIG_GET(context.Context, *CONFIG_MANAGE_CONFIG_GET_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_CONFIG_GET not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_CONFIG_IMPORT(context.Context, *CONFIG_MANAGE_CONFIG_IMPORT_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_CONFIG_IMPORT not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_CONFIG_POST(context.Context, *CONFIG_MANAGE_CONFIG_POST_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_CONFIG_POST not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_CONFIG_PUT(context.Context, *CONFIG_MANAGE_CONFIG_PUT_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_CONFIG_PUT not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_DEPLOY_CONFIG_GET(context.Context, *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_DEPLOY_CONFIG_GET not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_MULTI_NS_CONFIG_GET(context.Context, *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_MULTI_NS_CONFIG_GET not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_NAMESPACE_DEL(context.Context, *CONFIG_MANAGE_NAMESPACE_DEL_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_NAMESPACE_DEL not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_NAMESPACE_FIX(context.Context, *CONFIG_MANAGE_NAMESPACE_FIX_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_NAMESPACE_FIX not implemented")
}
func (*UnimplementedDopConfigManageServer) CONFIG_MANAGE_NAMESPACE_POST(context.Context, *CONFIG_MANAGE_NAMESPACE_POST_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CONFIG_MANAGE_NAMESPACE_POST not implemented")
}

func RegisterDopConfigManageServer(s grpc1.ServiceRegistrar, srv DopConfigManageServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_DopConfigManage_serviceDesc(srv, opts...), srv)
}

var _DopConfigManage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.openapiv1.dop.dop_config_manage",
	HandlerType: (*DopConfigManageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "dop_config_manage.proto",
}

func _get_DopConfigManage_serviceDesc(srv DopConfigManageServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_CONFIG_DEL(ctx, req.(*CONFIG_MANAGE_CONFIG_DEL_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_CONFIG_DEL", srv)
		_DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_CONFIG_EXPORT(ctx, req.(*CONFIG_MANAGE_CONFIG_EXPORT_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_CONFIG_EXPORT", srv)
		_DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_CONFIG_GET_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_CONFIG_GET(ctx, req.(*CONFIG_MANAGE_CONFIG_GET_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_CONFIG_GET_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_CONFIG_GET_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_CONFIG_GET", srv)
		_DopConfigManage_CONFIG_MANAGE_CONFIG_GET_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_CONFIG_GET_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_CONFIG_IMPORT(ctx, req.(*CONFIG_MANAGE_CONFIG_IMPORT_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_CONFIG_IMPORT", srv)
		_DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_CONFIG_POST_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_CONFIG_POST(ctx, req.(*CONFIG_MANAGE_CONFIG_POST_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_CONFIG_POST_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_CONFIG_POST_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_CONFIG_POST", srv)
		_DopConfigManage_CONFIG_MANAGE_CONFIG_POST_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_CONFIG_POST_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_CONFIG_PUT(ctx, req.(*CONFIG_MANAGE_CONFIG_PUT_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_CONFIG_PUT", srv)
		_DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_DEPLOY_CONFIG_GET(ctx, req.(*CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_DEPLOY_CONFIG_GET", srv)
		_DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_MULTI_NS_CONFIG_GET(ctx, req.(*CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_MULTI_NS_CONFIG_GET", srv)
		_DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_NAMESPACE_DEL(ctx, req.(*CONFIG_MANAGE_NAMESPACE_DEL_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_NAMESPACE_DEL", srv)
		_DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_NAMESPACE_FIX(ctx, req.(*CONFIG_MANAGE_NAMESPACE_FIX_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_NAMESPACE_FIX", srv)
		_DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_Handler)
	}

	_DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CONFIG_MANAGE_NAMESPACE_POST(ctx, req.(*CONFIG_MANAGE_NAMESPACE_POST_Request))
	}
	var _DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_info = transport.NewServiceInfo("erda.openapiv1.dop.dop_config_manage", "CONFIG_MANAGE_NAMESPACE_POST", srv)
		_DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_Handler = h.Interceptor(_DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_Handler)
	}

	var serviceDesc = _DopConfigManage_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CONFIG_MANAGE_CONFIG_DEL",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_CONFIG_DEL_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_CONFIG_DEL(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_DEL",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_CONFIG_DEL_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_CONFIG_EXPORT",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_CONFIG_EXPORT_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_CONFIG_EXPORT(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_EXPORT",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_CONFIG_EXPORT_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_CONFIG_GET",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_CONFIG_GET_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_CONFIG_GET(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_CONFIG_GET_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_CONFIG_GET_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_GET",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_CONFIG_GET_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_CONFIG_IMPORT",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_CONFIG_IMPORT_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_CONFIG_IMPORT(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_IMPORT",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_CONFIG_IMPORT_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_CONFIG_POST",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_CONFIG_POST_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_CONFIG_POST(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_CONFIG_POST_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_CONFIG_POST_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_POST",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_CONFIG_POST_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_CONFIG_PUT",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_CONFIG_PUT_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_CONFIG_PUT(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_CONFIG_PUT",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_CONFIG_PUT_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_DEPLOY_CONFIG_GET",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_DEPLOY_CONFIG_GET(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_DEPLOY_CONFIG_GET",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_DEPLOY_CONFIG_GET_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_MULTI_NS_CONFIG_GET",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_MULTI_NS_CONFIG_GET(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_MULTI_NS_CONFIG_GET",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_NAMESPACE_DEL",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_NAMESPACE_DEL_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_NAMESPACE_DEL(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_NAMESPACE_DEL",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_NAMESPACE_DEL_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_NAMESPACE_FIX",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_NAMESPACE_FIX_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_NAMESPACE_FIX(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_NAMESPACE_FIX",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_NAMESPACE_FIX_Handler)
			},
		},
		{
			MethodName: "CONFIG_MANAGE_NAMESPACE_POST",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CONFIG_MANAGE_NAMESPACE_POST_Request)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DopConfigManageServer).CONFIG_MANAGE_NAMESPACE_POST(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_info)
				}
				if interceptor == nil {
					return _DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.openapiv1.dop.dop_config_manage/CONFIG_MANAGE_NAMESPACE_POST",
				}
				return interceptor(ctx, in, info, _DopConfigManage_CONFIG_MANAGE_NAMESPACE_POST_Handler)
			},
		},
	}
	return &serviceDesc
}
