// Package json
// @title: client
// @description: 描述
// @author: CaoChao
// @date: 2024-05-02
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接错误")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err = client.Call("HelloService.Hello", "CaoChao", &reply)
	if err != nil {
		panic("调用错误")
	}
	fmt.Println(reply)
}
