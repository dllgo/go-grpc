package grpcplus

import (
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// NewClient returns a new go-micro clinet
func NewClient(opts ...client.Option) client.Client {
	//etcd registry
	r := etcd.NewRegistry(etcdOptions)

	// set the registry and selector
	options := []client.Option{
		client.Registry(r),
	}
	// append user options
	options = append(options, opts...)
	// return a micro.Client
	return grpc.NewClient(options...)
}
