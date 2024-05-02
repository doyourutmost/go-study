// Package http_server
// @title: client
// @description: 描述
// @author: CaoChao
// @date: 2024-05-02
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var helloReply string
	if err = client.Call("HelloService.Hello", "world~~", &helloReply); err != nil {
		log.Fatal(err)
	}

	fmt.Println(helloReply)
}
