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
	"learn-go/rpc/stub/server_proxy"
)

func main() {
	// 1. 实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑 handler
	_ = server_proxy.RegisterHelloService(new(handler.HelloService))
	// 3. 启动服务
	for {
		conn, _ := listener.Accept() // 阻塞等待客户端连接
		go rpc.ServeConn(conn)
	}
}
