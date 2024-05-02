// Package json
// @title: server
// @description: 描述
// @author: CaoChao
// @date: 2024-05-02

package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("启动错误")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("建立链接失败")
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
