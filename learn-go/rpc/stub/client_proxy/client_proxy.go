// Package client_proxy
// @title: client_proxy
// @description: 描述
// @author: CaoChao
// @date: 2024-05-03
package client_proxy

import (
	"net/rpc"

	"learn-go/rpc/stub/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) *HelloServiceStub {
	client, err := rpc.Dial(protocol, address)
	if err != nil {
		panic(err)
	}
	return &HelloServiceStub{Client: client}
}

func (c *HelloServiceStub) Hello(request string, reply *string) (err error) {
	err = c.Call(handler.HelloServiceName+".Hello", request, reply)
	return
}
