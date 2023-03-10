// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: search.proto

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewSearchServiceClient(cc grpc1.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.search.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations should embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedSearchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterSearchServiceServer(s grpc1.ServiceRegistrar, srv SearchServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_SearchService_serviceDesc(srv, opts...), srv)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.dop.search.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "search.proto",
}

func _get_SearchService_serviceDesc(srv SearchServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_SearchService_Search_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Search(ctx, req.(*SearchRequest))
	}
	var _SearchService_Search_info transport.ServiceInfo
	if h.Interceptor != nil {
		_SearchService_Search_info = transport.NewServiceInfo("erda.dop.search.SearchService", "Search", srv)
		_SearchService_Search_Handler = h.Interceptor(_SearchService_Search_Handler)
	}

	var serviceDesc = _SearchService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(SearchRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(SearchServiceServer).Search(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _SearchService_Search_info)
				}
				if interceptor == nil {
					return _SearchService_Search_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.search.SearchService/Search",
				}
				return interceptor(ctx, in, info, _SearchService_Search_Handler)
			},
		},
	}
	return &serviceDesc
}
