// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: source.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterSourceServiceImp source.proto
func RegisterSourceServiceImp(regester transport.Register, srv SourceServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterSourceServiceHandler(regester, SourceServiceHandler(srv), _ops.HTTP...)
	RegisterSourceServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.source.SourceService",
	)
}

var (
	sourceServiceClientType  = reflect.TypeOf((*SourceServiceClient)(nil)).Elem()
	sourceServiceServerType  = reflect.TypeOf((*SourceServiceServer)(nil)).Elem()
	sourceServiceHandlerType = reflect.TypeOf((*SourceServiceHandler)(nil)).Elem()
)

// SourceServiceClientType .
func SourceServiceClientType() reflect.Type { return sourceServiceClientType }

// SourceServiceServerType .
func SourceServiceServerType() reflect.Type { return sourceServiceServerType }

// SourceServiceHandlerType .
func SourceServiceHandlerType() reflect.Type { return sourceServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		sourceServiceClientType,
		// server types
		sourceServiceServerType,
		// handler types
		sourceServiceHandlerType,
	}
}
