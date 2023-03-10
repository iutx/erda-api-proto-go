// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: msp_tenant_project.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterMspTenantProjectImp msp_tenant_project.proto
func RegisterMspTenantProjectImp(regester transport.Register, srv MspTenantProjectServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterMspTenantProjectHandler(regester, MspTenantProjectHandler(srv), _ops.HTTP...)
	RegisterMspTenantProjectServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.openapiv1.msp.MspTenantProject",
	)
}

var (
	mspTenantProjectClientType  = reflect.TypeOf((*MspTenantProjectClient)(nil)).Elem()
	mspTenantProjectServerType  = reflect.TypeOf((*MspTenantProjectServer)(nil)).Elem()
	mspTenantProjectHandlerType = reflect.TypeOf((*MspTenantProjectHandler)(nil)).Elem()
)

// MspTenantProjectClientType .
func MspTenantProjectClientType() reflect.Type { return mspTenantProjectClientType }

// MspTenantProjectServerType .
func MspTenantProjectServerType() reflect.Type { return mspTenantProjectServerType }

// MspTenantProjectHandlerType .
func MspTenantProjectHandlerType() reflect.Type { return mspTenantProjectHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		mspTenantProjectClientType,
		// server types
		mspTenantProjectServerType,
		// handler types
		mspTenantProjectHandlerType,
	}
}
