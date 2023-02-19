package services

import (
	"context"
	"fmt"
	pb "gorpc/libs/rpc/src"
)

type HelloServer struct {
}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("a new SayHello request")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
