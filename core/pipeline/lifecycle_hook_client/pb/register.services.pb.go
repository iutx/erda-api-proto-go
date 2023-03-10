// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: lifecycle_hook_client.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterLifeCycleServiceImp lifecycle_hook_client.proto
func RegisterLifeCycleServiceImp(regester transport.Register, srv LifeCycleServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterLifeCycleServiceHandler(regester, LifeCycleServiceHandler(srv), _ops.HTTP...)
	RegisterLifeCycleServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.lifecycle_hook_client.LifeCycleService",
	)
}

var (
	lifeCycleServiceClientType  = reflect.TypeOf((*LifeCycleServiceClient)(nil)).Elem()
	lifeCycleServiceServerType  = reflect.TypeOf((*LifeCycleServiceServer)(nil)).Elem()
	lifeCycleServiceHandlerType = reflect.TypeOf((*LifeCycleServiceHandler)(nil)).Elem()
)

// LifeCycleServiceClientType .
func LifeCycleServiceClientType() reflect.Type { return lifeCycleServiceClientType }

// LifeCycleServiceServerType .
func LifeCycleServiceServerType() reflect.Type { return lifeCycleServiceServerType }

// LifeCycleServiceHandlerType .
func LifeCycleServiceHandlerType() reflect.Type { return lifeCycleServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		lifeCycleServiceClientType,
		// server types
		lifeCycleServiceServerType,
		// handler types
		lifeCycleServiceHandlerType,
	}
}
