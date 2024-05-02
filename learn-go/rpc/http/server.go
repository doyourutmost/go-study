// Package http
// @title: server
// @description: 描述
// @author: CaoChao
// @date: 2024-05-02
package main

import (
	"io"
	"net/http"
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
	http.HandleFunc("/json-rpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic("启动错误")
	}
}
