// Package server
// @title: server
// @description: 描述
// @author: CaoChao
// @date: 2024-05-03
package main

import (
	"net"
	"net/rpc"

	"learn-go/rpc/stub/handler"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) (err error) {
	*reply = "hello " + request
	return
}

func main() {
	// 1. 实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑 handler
	_ = rpc.RegisterName(handler.HelloServiceName, &HelloService{})
	// 3. 启动服务
	for {
		conn, _ := listener.Accept() // 阻塞等待客户端连接
		go rpc.ServeConn(conn)
	}
}
