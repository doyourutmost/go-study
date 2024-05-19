package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"learn-go/rpc/grpc/protobuf/proto"
)

var (
	addr = flag.String("addr", "localhost:8051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}(conn)

	c := proto.NewEchoClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	address := &proto.Address{
		Province: "HuNan",
		City:     "ChangSha",
	}
	companyAddress := &proto.Address{
		Province: "GuangDong",
		City:     "ShenZhen",
	}
	company := &proto.EchoRequest_Company{
		Name:    "Tencent",
		Address: companyAddress,
	}
	m := make(map[string]string)
	m["key"] = "value"
	m["key2"] = "value2"
	email := "1234567@qq.com"

	a, err := anypb.New(address)
	timestamp := timestamppb.New(time.Now())

	echo, err := c.TestEcho(ctx, &proto.EchoRequest{
		Name:      "CaoChao",
		Hobbies:   []string{"basketball", "football"},
		Email:     &email,
		Sex:       proto.Sex_OTHER,
		Address:   address,
		Company:   company,
		Any:       a,
		Map:       m,
		Timestamp: timestamp,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Receive: %v", echo)
}
