// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: rule.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/rule/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.dop.rule",
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
	ruleServiceClientType = reflect.TypeOf((*pb.RuleServiceClient)(nil)).Elem()
	ruleServiceServerType = reflect.TypeOf((*pb.RuleServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.dop.rule-client":
		return p.client
	case "erda.dop.rule.RuleService":
		return &ruleServiceWrapper{client: p.client.RuleService(), opts: opts}
	case "erda.dop.rule.RuleService.client":
		return p.client.RuleService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case ruleServiceClientType:
		return p.client.RuleService()
	case ruleServiceServerType:
		return &ruleServiceWrapper{client: p.client.RuleService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.rule-client", &servicehub.Spec{
		Services: []string{
			"erda.dop.rule.RuleService",
			"erda.dop.rule.RuleService.client",
			"erda.dop.rule-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			ruleServiceClientType,
			// server types
			ruleServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
