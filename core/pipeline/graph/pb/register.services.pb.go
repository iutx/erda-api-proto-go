// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: graph.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterGraphServiceImp graph.proto
func RegisterGraphServiceImp(regester transport.Register, srv GraphServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterGraphServiceHandler(regester, GraphServiceHandler(srv), _ops.HTTP...)
	RegisterGraphServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.graph.GraphService",
	)
}

var (
	graphServiceClientType  = reflect.TypeOf((*GraphServiceClient)(nil)).Elem()
	graphServiceServerType  = reflect.TypeOf((*GraphServiceServer)(nil)).Elem()
	graphServiceHandlerType = reflect.TypeOf((*GraphServiceHandler)(nil)).Elem()
)

// GraphServiceClientType .
func GraphServiceClientType() reflect.Type { return graphServiceClientType }

// GraphServiceServerType .
func GraphServiceServerType() reflect.Type { return graphServiceServerType }

// GraphServiceHandlerType .
func GraphServiceHandlerType() reflect.Type { return graphServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		graphServiceClientType,
		// server types
		graphServiceServerType,
		// handler types
		graphServiceHandlerType,
	}
}
