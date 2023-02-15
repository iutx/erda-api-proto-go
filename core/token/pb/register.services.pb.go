// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: token.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterTokenServiceImp token.proto
func RegisterTokenServiceImp(regester transport.Register, srv TokenServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterTokenServiceHandler(regester, TokenServiceHandler(srv), _ops.HTTP...)
	RegisterTokenServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.token.TokenService",
	)
}

var (
	tokenServiceClientType  = reflect.TypeOf((*TokenServiceClient)(nil)).Elem()
	tokenServiceServerType  = reflect.TypeOf((*TokenServiceServer)(nil)).Elem()
	tokenServiceHandlerType = reflect.TypeOf((*TokenServiceHandler)(nil)).Elem()
)

// TokenServiceClientType .
func TokenServiceClientType() reflect.Type { return tokenServiceClientType }

// TokenServiceServerType .
func TokenServiceServerType() reflect.Type { return tokenServiceServerType }

// TokenServiceHandlerType .
func TokenServiceHandlerType() reflect.Type { return tokenServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		tokenServiceClientType,
		// server types
		tokenServiceServerType,
		// handler types
		tokenServiceHandlerType,
	}
}
