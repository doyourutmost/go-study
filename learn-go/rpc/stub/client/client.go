// Package client
// @title: client
// @description: 描述
// @author: CaoChao
// @date: 2024-05-03
package main

import (
	"fmt"

	"learn-go/rpc/stub/client_proxy"
)

func main() {
	// 1. 建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string
	err := client.Hello("world 111", &reply)
	if err != nil {
		panic(fmt.Sprintf("调用失败: %s", err))
	}
	fmt.Println(reply)
}
