// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: publishitem.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/publishitem/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.dop.publishitem",
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
	clientsType                  = reflect.TypeOf((*Client)(nil)).Elem()
	publishItemServiceClientType = reflect.TypeOf((*pb.PublishItemServiceClient)(nil)).Elem()
	publishItemServiceServerType = reflect.TypeOf((*pb.PublishItemServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.dop.publishitem-client":
		return p.client
	case "erda.dop.publishitem.PublishItemService":
		return &publishItemServiceWrapper{client: p.client.PublishItemService(), opts: opts}
	case "erda.dop.publishitem.PublishItemService.client":
		return p.client.PublishItemService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case publishItemServiceClientType:
		return p.client.PublishItemService()
	case publishItemServiceServerType:
		return &publishItemServiceWrapper{client: p.client.PublishItemService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.publishitem-client", &servicehub.Spec{
		Services: []string{
			"erda.dop.publishitem.PublishItemService",
			"erda.dop.publishitem.PublishItemService.client",
			"erda.dop.publishitem-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			publishItemServiceClientType,
			// server types
			publishItemServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
