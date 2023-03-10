// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: msp_apm_log_service.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterMspApmLogServiceImp msp_apm_log_service.proto
func RegisterMspApmLogServiceImp(regester transport.Register, srv MspApmLogServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterMspApmLogServiceHandler(regester, MspApmLogServiceHandler(srv), _ops.HTTP...)
	RegisterMspApmLogServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.openapiv1.msp.MspApmLogService",
	)
}

var (
	mspApmLogServiceClientType  = reflect.TypeOf((*MspApmLogServiceClient)(nil)).Elem()
	mspApmLogServiceServerType  = reflect.TypeOf((*MspApmLogServiceServer)(nil)).Elem()
	mspApmLogServiceHandlerType = reflect.TypeOf((*MspApmLogServiceHandler)(nil)).Elem()
)

// MspApmLogServiceClientType .
func MspApmLogServiceClientType() reflect.Type { return mspApmLogServiceClientType }

// MspApmLogServiceServerType .
func MspApmLogServiceServerType() reflect.Type { return mspApmLogServiceServerType }

// MspApmLogServiceHandlerType .
func MspApmLogServiceHandlerType() reflect.Type { return mspApmLogServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		mspApmLogServiceClientType,
		// server types
		mspApmLogServiceServerType,
		// handler types
		mspApmLogServiceHandlerType,
	}
}
