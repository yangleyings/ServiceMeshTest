// Code generated by Kitex v0.0.1. DO NOT EDIT.
package hello

import (
	"github.com/cloudwego/kitex/server"
	pb "github.com/rpcxio/rpcx-benchmark/proto"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler pb.Hello, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
