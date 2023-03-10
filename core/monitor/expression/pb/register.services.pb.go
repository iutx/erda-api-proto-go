// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: expression.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterExpressionServiceImp expression.proto
func RegisterExpressionServiceImp(regester transport.Register, srv ExpressionServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterExpressionServiceHandler(regester, ExpressionServiceHandler(srv), _ops.HTTP...)
	RegisterExpressionServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.monitor.expression.ExpressionService",
	)
}

var (
	expressionServiceClientType  = reflect.TypeOf((*ExpressionServiceClient)(nil)).Elem()
	expressionServiceServerType  = reflect.TypeOf((*ExpressionServiceServer)(nil)).Elem()
	expressionServiceHandlerType = reflect.TypeOf((*ExpressionServiceHandler)(nil)).Elem()
)

// ExpressionServiceClientType .
func ExpressionServiceClientType() reflect.Type { return expressionServiceClientType }

// ExpressionServiceServerType .
func ExpressionServiceServerType() reflect.Type { return expressionServiceServerType }

// ExpressionServiceHandlerType .
func ExpressionServiceHandlerType() reflect.Type { return expressionServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		expressionServiceClientType,
		// server types
		expressionServiceServerType,
		// handler types
		expressionServiceHandlerType,
	}
}
