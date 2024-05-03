// Package handler
// @title: handler
// @description: 描述
// @author: CaoChao
// @date: 2024-05-03
package handler

const (
	// HelloServiceName 加个前缀，避免和其他服务冲突
	HelloServiceName = "handler/HelloService"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}
