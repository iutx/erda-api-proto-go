// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: publishitem.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterPublishItemServiceImp publishitem.proto
func RegisterPublishItemServiceImp(regester transport.Register, srv PublishItemServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterPublishItemServiceHandler(regester, PublishItemServiceHandler(srv), _ops.HTTP...)
	RegisterPublishItemServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.dop.publishitem.PublishItemService",
	)
}

var (
	publishItemServiceClientType  = reflect.TypeOf((*PublishItemServiceClient)(nil)).Elem()
	publishItemServiceServerType  = reflect.TypeOf((*PublishItemServiceServer)(nil)).Elem()
	publishItemServiceHandlerType = reflect.TypeOf((*PublishItemServiceHandler)(nil)).Elem()
)

// PublishItemServiceClientType .
func PublishItemServiceClientType() reflect.Type { return publishItemServiceClientType }

// PublishItemServiceServerType .
func PublishItemServiceServerType() reflect.Type { return publishItemServiceServerType }

// PublishItemServiceHandlerType .
func PublishItemServiceHandlerType() reflect.Type { return publishItemServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		publishItemServiceClientType,
		// server types
		publishItemServiceServerType,
		// handler types
		publishItemServiceHandlerType,
	}
}
