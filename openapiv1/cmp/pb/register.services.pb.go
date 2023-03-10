// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: cmp.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterCmpImp cmp.proto
func RegisterCmpImp(regester transport.Register, srv CmpServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterCmpHandler(regester, CmpHandler(srv), _ops.HTTP...)
	RegisterCmpServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.openapiv1.cmp.Cmp",
	)
}

var (
	cmpClientType  = reflect.TypeOf((*CmpClient)(nil)).Elem()
	cmpServerType  = reflect.TypeOf((*CmpServer)(nil)).Elem()
	cmpHandlerType = reflect.TypeOf((*CmpHandler)(nil)).Elem()
)

// CmpClientType .
func CmpClientType() reflect.Type { return cmpClientType }

// CmpServerType .
func CmpServerType() reflect.Type { return cmpServerType }

// CmpHandlerType .
func CmpHandlerType() reflect.Type { return cmpHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		cmpClientType,
		// server types
		cmpServerType,
		// handler types
		cmpHandlerType,
	}
}
