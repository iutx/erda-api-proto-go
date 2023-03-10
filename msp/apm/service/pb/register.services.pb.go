// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: service.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterApmServiceServiceImp service.proto
func RegisterApmServiceServiceImp(regester transport.Register, srv ApmServiceServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterApmServiceServiceHandler(regester, ApmServiceServiceHandler(srv), _ops.HTTP...)
	RegisterApmServiceServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.msp.apm.service.ApmServiceService",
	)
}

var (
	apmServiceServiceClientType  = reflect.TypeOf((*ApmServiceServiceClient)(nil)).Elem()
	apmServiceServiceServerType  = reflect.TypeOf((*ApmServiceServiceServer)(nil)).Elem()
	apmServiceServiceHandlerType = reflect.TypeOf((*ApmServiceServiceHandler)(nil)).Elem()
)

// ApmServiceServiceClientType .
func ApmServiceServiceClientType() reflect.Type { return apmServiceServiceClientType }

// ApmServiceServiceServerType .
func ApmServiceServiceServerType() reflect.Type { return apmServiceServiceServerType }

// ApmServiceServiceHandlerType .
func ApmServiceServiceHandlerType() reflect.Type { return apmServiceServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		apmServiceServiceClientType,
		// server types
		apmServiceServiceServerType,
		// handler types
		apmServiceServiceHandlerType,
	}
}
