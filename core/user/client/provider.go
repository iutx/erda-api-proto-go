// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: user.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/user/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.user",
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
	userServiceClientType = reflect.TypeOf((*pb.UserServiceClient)(nil)).Elem()
	userServiceServerType = reflect.TypeOf((*pb.UserServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.user-client":
		return p.client
	case "erda.core.user.UserService":
		return &userServiceWrapper{client: p.client.UserService(), opts: opts}
	case "erda.core.user.UserService.client":
		return p.client.UserService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case userServiceClientType:
		return p.client.UserService()
	case userServiceServerType:
		return &userServiceWrapper{client: p.client.UserService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.user-client", &servicehub.Spec{
		Services: []string{
			"erda.core.user.UserService",
			"erda.core.user.UserService.client",
			"erda.core.user-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			userServiceClientType,
			// server types
			userServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}