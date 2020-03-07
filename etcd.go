package grpcplus

import (
	"time"

	"github.com/micro/go-micro/v2/registry"
)

func etcdOptions(op *registry.Options) {
	op.Addrs = []string{
		"http://127.0.0.1:2379",
	}
	op.Timeout = 5 * time.Second
}
