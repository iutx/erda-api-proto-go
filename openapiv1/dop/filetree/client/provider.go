// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: dop_filetree.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/dop/filetree/pb"
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
	clientsType           = reflect.TypeOf((*Client)(nil)).Elem()
	dopFiletreeClientType = reflect.TypeOf((*pb.DopFiletreeClient)(nil)).Elem()
	dopFiletreeServerType = reflect.TypeOf((*pb.DopFiletreeServer)(nil)).Elem()
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
	case "erda.openapiv1.dop.DopFiletree":
		return &dopFiletreeWrapper{client: p.client.DopFiletree(), opts: opts}
	case "erda.openapiv1.dop.DopFiletree.client":
		return p.client.DopFiletree()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case dopFiletreeClientType:
		return p.client.DopFiletree()
	case dopFiletreeServerType:
		return &dopFiletreeWrapper{client: p.client.DopFiletree(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.openapiv1.dop-client", &servicehub.Spec{
		Services: []string{
			"erda.openapiv1.dop.DopFiletree",
			"erda.openapiv1.dop.DopFiletree.client",
			"erda.openapiv1.dop-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			dopFiletreeClientType,
			// server types
			dopFiletreeServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
