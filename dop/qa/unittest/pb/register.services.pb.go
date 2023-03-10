// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: unittest.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterUnitTestServiceImp unittest.proto
func RegisterUnitTestServiceImp(regester transport.Register, srv UnitTestServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterUnitTestServiceHandler(regester, UnitTestServiceHandler(srv), _ops.HTTP...)
	RegisterUnitTestServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.dop.qa.unittest.UnitTestService",
	)
}

var (
	unitTestServiceClientType  = reflect.TypeOf((*UnitTestServiceClient)(nil)).Elem()
	unitTestServiceServerType  = reflect.TypeOf((*UnitTestServiceServer)(nil)).Elem()
	unitTestServiceHandlerType = reflect.TypeOf((*UnitTestServiceHandler)(nil)).Elem()
)

// UnitTestServiceClientType .
func UnitTestServiceClientType() reflect.Type { return unitTestServiceClientType }

// UnitTestServiceServerType .
func UnitTestServiceServerType() reflect.Type { return unitTestServiceServerType }

// UnitTestServiceHandlerType .
func UnitTestServiceHandlerType() reflect.Type { return unitTestServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		unitTestServiceClientType,
		// server types
		unitTestServiceServerType,
		// handler types
		unitTestServiceHandlerType,
	}
}
