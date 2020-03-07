package grpcplus

import (
	"time"

	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
)

// NewService returns a new go-micro service
func NewService(opts ...service.Option) service.Service {
	//etcd registry
	reg := etcd.NewRegistry(etcdOptions)

	// set the registry and selector
	options := []service.Option{
		service.Registry(reg),
		service.RegisterTTL(time.Second * 15),      //重新注册时间
		service.RegisterInterval(time.Second * 10), //注册过期时间
		service.Version("latest"),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpc.NewService(options...)
}
