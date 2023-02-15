// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: api_gateway.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterApiGatewayImp api_gateway.proto
func RegisterApiGatewayImp(regester transport.Register, srv ApiGatewayServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterApiGatewayHandler(regester, ApiGatewayHandler(srv), _ops.HTTP...)
	RegisterApiGatewayServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.openapiv1.api.ApiGateway",
	)
}

var (
	apiGatewayClientType  = reflect.TypeOf((*ApiGatewayClient)(nil)).Elem()
	apiGatewayServerType  = reflect.TypeOf((*ApiGatewayServer)(nil)).Elem()
	apiGatewayHandlerType = reflect.TypeOf((*ApiGatewayHandler)(nil)).Elem()
)

// ApiGatewayClientType .
func ApiGatewayClientType() reflect.Type { return apiGatewayClientType }

// ApiGatewayServerType .
func ApiGatewayServerType() reflect.Type { return apiGatewayServerType }

// ApiGatewayHandlerType .
func ApiGatewayHandlerType() reflect.Type { return apiGatewayHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		apiGatewayClientType,
		// server types
		apiGatewayServerType,
		// handler types
		apiGatewayHandlerType,
	}
}