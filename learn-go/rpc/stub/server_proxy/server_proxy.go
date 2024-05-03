// Package server_proxy
// @title: server_proxy
// @description: 描述
// @author: CaoChao
// @date: 2024-05-03
package server_proxy

import (
	"net/rpc"

	"learn-go/rpc/stub/handler"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
