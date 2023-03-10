// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: fdp.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/fdp/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.openapiv1.fdp",
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
	clientsType   = reflect.TypeOf((*Client)(nil)).Elem()
	fdpClientType = reflect.TypeOf((*pb.FdpClient)(nil)).Elem()
	fdpServerType = reflect.TypeOf((*pb.FdpServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.openapiv1.fdp-client":
		return p.client
	case "erda.openapiv1.fdp.Fdp":
		return &fdpWrapper{client: p.client.Fdp(), opts: opts}
	case "erda.openapiv1.fdp.Fdp.client":
		return p.client.Fdp()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case fdpClientType:
		return p.client.Fdp()
	case fdpServerType:
		return &fdpWrapper{client: p.client.Fdp(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.openapiv1.fdp-client", &servicehub.Spec{
		Services: []string{
			"erda.openapiv1.fdp.Fdp",
			"erda.openapiv1.fdp.Fdp.client",
			"erda.openapiv1.fdp-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			fdpClientType,
			// server types
			fdpServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
