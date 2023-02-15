// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: elf.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/elf/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.openapiv1.elf",
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
	elfClientType = reflect.TypeOf((*pb.ElfClient)(nil)).Elem()
	elfServerType = reflect.TypeOf((*pb.ElfServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.openapiv1.elf-client":
		return p.client
	case "erda.openapiv1.elf.Elf":
		return &elfWrapper{client: p.client.Elf(), opts: opts}
	case "erda.openapiv1.elf.Elf.client":
		return p.client.Elf()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case elfClientType:
		return p.client.Elf()
	case elfServerType:
		return &elfWrapper{client: p.client.Elf(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.openapiv1.elf-client", &servicehub.Spec{
		Services: []string{
			"erda.openapiv1.elf.Elf",
			"erda.openapiv1.elf.Elf.client",
			"erda.openapiv1.elf-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			elfClientType,
			// server types
			elfServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
