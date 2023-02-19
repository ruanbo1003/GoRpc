package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pbHello "gorpc/libs/rpc/src"
	"log"
	"time"
)

const (
	serverAddress  = "localhost:8801"
	defaultName    = "world"
	requestTimeOut = time.Second * 10
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc client connect to server failed:%v", err)
	}
	defer conn.Close()

	c := pbHello.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)
	defer cancel()

	rsp, err := c.SayHello(ctx, &pbHello.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("client say hello error: %v", err)
	}

	log.Println("say hello response:", rsp.Message)

	log.Println("stream-client")
	{
		streamClient := StreamClient{}
		streamClient.ReqLists(conn, ctx)

		streamClient.ReqRecord(conn, ctx)

		streamClient.ReqRoute(conn, ctx)
	}

	fmt.Println("all request sent")
	time.Sleep(time.Second * 2)
}
