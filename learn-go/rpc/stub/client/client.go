// Package client
// @title: client
// @description: 描述
// @author: CaoChao
// @date: 2024-05-03
package main

import (
	"fmt"
	"net/rpc"

	"learn-go/rpc/stub/handler"
)

func main() {
	// 1. 建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call(handler.HelloServiceName+".Hello", "world~~~", &reply)
	if err != nil {
		panic(fmt.Sprintf("调用失败: %s", err))
	}
	fmt.Println(reply)
}
