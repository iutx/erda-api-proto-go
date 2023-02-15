// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: uc.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterUcImp uc.proto
func RegisterUcImp(regester transport.Register, srv UcServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterUcHandler(regester, UcHandler(srv), _ops.HTTP...)
	RegisterUcServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.openapiv1.uc.Uc",
	)
}

var (
	ucClientType  = reflect.TypeOf((*UcClient)(nil)).Elem()
	ucServerType  = reflect.TypeOf((*UcServer)(nil)).Elem()
	ucHandlerType = reflect.TypeOf((*UcHandler)(nil)).Elem()
)

// UcClientType .
func UcClientType() reflect.Type { return ucClientType }

// UcServerType .
func UcServerType() reflect.Type { return ucServerType }

// UcHandlerType .
func UcHandlerType() reflect.Type { return ucHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		ucClientType,
		// server types
		ucServerType,
		// handler types
		ucHandlerType,
	}
}
