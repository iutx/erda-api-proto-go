// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: rule.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterRuleServiceImp rule.proto
func RegisterRuleServiceImp(regester transport.Register, srv RuleServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterRuleServiceHandler(regester, RuleServiceHandler(srv), _ops.HTTP...)
	RegisterRuleServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.dop.rule.RuleService",
	)
}

var (
	ruleServiceClientType  = reflect.TypeOf((*RuleServiceClient)(nil)).Elem()
	ruleServiceServerType  = reflect.TypeOf((*RuleServiceServer)(nil)).Elem()
	ruleServiceHandlerType = reflect.TypeOf((*RuleServiceHandler)(nil)).Elem()
)

// RuleServiceClientType .
func RuleServiceClientType() reflect.Type { return ruleServiceClientType }

// RuleServiceServerType .
func RuleServiceServerType() reflect.Type { return ruleServiceServerType }

// RuleServiceHandlerType .
func RuleServiceHandlerType() reflect.Type { return ruleServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		ruleServiceClientType,
		// server types
		ruleServiceServerType,
		// handler types
		ruleServiceHandlerType,
	}
}
