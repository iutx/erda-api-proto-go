// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: dop_config_manage.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/dop/config/manage/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.openapiv1.dop",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType               = reflect.TypeOf((*Client)(nil)).Elem()
	dopConfigManageClientType = reflect.TypeOf((*pb.DopConfigManageClient)(nil)).Elem()
	dopConfigManageServerType = reflect.TypeOf((*pb.DopConfigManageServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.openapiv1.dop-client":
		return p.client
	case "erda.openapiv1.dop.DopConfigManage":
		return &dopConfigManageWrapper{client: p.client.DopConfigManage(), opts: opts}
	case "erda.openapiv1.dop.DopConfigManage.client":
		return p.client.DopConfigManage()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case dopConfigManageClientType:
		return p.client.DopConfigManage()
	case dopConfigManageServerType:
		return &dopConfigManageWrapper{client: p.client.DopConfigManage(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.openapiv1.dop-client", &servicehub.Spec{
		Services: []string{
			"erda.openapiv1.dop.DopConfigManage",
			"erda.openapiv1.dop.DopConfigManage.client",
			"erda.openapiv1.dop-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			dopConfigManageClientType,
			// server types
			dopConfigManageServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
