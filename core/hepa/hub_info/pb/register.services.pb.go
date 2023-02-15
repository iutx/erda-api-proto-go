// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: hub_info.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterHubInfoServiceImp hub_info.proto
func RegisterHubInfoServiceImp(regester transport.Register, srv HubInfoServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterHubInfoServiceHandler(regester, HubInfoServiceHandler(srv), _ops.HTTP...)
	RegisterHubInfoServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.hepa.hub_info.HubInfoService",
	)
}

var (
	hubInfoServiceClientType  = reflect.TypeOf((*HubInfoServiceClient)(nil)).Elem()
	hubInfoServiceServerType  = reflect.TypeOf((*HubInfoServiceServer)(nil)).Elem()
	hubInfoServiceHandlerType = reflect.TypeOf((*HubInfoServiceHandler)(nil)).Elem()
)

// HubInfoServiceClientType .
func HubInfoServiceClientType() reflect.Type { return hubInfoServiceClientType }

// HubInfoServiceServerType .
func HubInfoServiceServerType() reflect.Type { return hubInfoServiceServerType }

// HubInfoServiceHandlerType .
func HubInfoServiceHandlerType() reflect.Type { return hubInfoServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		hubInfoServiceClientType,
		// server types
		hubInfoServiceServerType,
		// handler types
		hubInfoServiceHandlerType,
	}
}
