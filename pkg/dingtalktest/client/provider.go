// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: dingtalktest.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/pkg/dingtalktest/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.pkg.dingtalktest",
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
	clientsType                   = reflect.TypeOf((*Client)(nil)).Elem()
	dingTalkTestServiceClientType = reflect.TypeOf((*pb.DingTalkTestServiceClient)(nil)).Elem()
	dingTalkTestServiceServerType = reflect.TypeOf((*pb.DingTalkTestServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.pkg.dingtalktest-client":
		return p.client
	case "erda.pkg.dingtalktest.DingTalkTestService":
		return &dingTalkTestServiceWrapper{client: p.client.DingTalkTestService(), opts: opts}
	case "erda.pkg.dingtalktest.DingTalkTestService.client":
		return p.client.DingTalkTestService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case dingTalkTestServiceClientType:
		return p.client.DingTalkTestService()
	case dingTalkTestServiceServerType:
		return &dingTalkTestServiceWrapper{client: p.client.DingTalkTestService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.pkg.dingtalktest-client", &servicehub.Spec{
		Services: []string{
			"erda.pkg.dingtalktest.DingTalkTestService",
			"erda.pkg.dingtalktest.DingTalkTestService.client",
			"erda.pkg.dingtalktest-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			dingTalkTestServiceClientType,
			// server types
			dingTalkTestServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
