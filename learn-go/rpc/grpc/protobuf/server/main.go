package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"learn-go/rpc/grpc/protobuf/proto"
)

var (
	port = flag.Int("port", 8051, "The server port")
)

type server struct {
	proto.UnimplementedEchoServer
}

func (s *server) TestEcho(ctx context.Context, in *proto.EchoRequest) (result *proto.EchoResponse, err error) {
	result = &proto.EchoResponse{}
	result.Name = in.Name
	result.Hobbies = in.Hobbies
	result.Email = in.Email
	result.Sex = in.Sex
	result.Address = in.Address
	company := &proto.EchoResponse_Company{
		Name:    in.GetCompany().GetName(),
		Address: in.GetCompany().GetAddress(),
	}
	result.Company = company
	result.Map = in.Map
	if in.Any.MessageIs(&proto.Address{}) {
		address := &proto.Address{}
		_ = in.Any.UnmarshalTo(address)
		fmt.Printf("Any 解压address: %v\n", address)

	}
	result.Any = in.Any
	result.Timestamp = in.Timestamp
	return
}

func (s *server) TestEmpty(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
