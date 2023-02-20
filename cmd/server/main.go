package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	services "gorpc/cmd/server/services"
	pb "gorpc/libs/rpc/src"
	"gorpc/libs/rpc/util"
	"log"
	"net"
	"time"
)

type serverConfig struct {
	Host         string        `json:",default=0.0.0.0"`
	Port         int           `json:",range=[80,65535)"`
	LogMode      string        `json:",options=[file,console]"`
	Verbose      bool          `json:",optional"`
	MaxConns     int           `json:",default=10000"`
	Timeout      time.Duration `json:",default=3s"`
	CpuThreshold int64         `json:",default=900,range=[0,1000]"`
}

func main() {
	listen, err := net.Listen("tcp", ":8801")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	fmt.Println("server listen on port:8001")

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			util.UnaryLoggingInterceptor),
		grpc_middleware.WithStreamServerChain(
			util.StreamLoggingInterceptor),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, &services.HelloServer{})
	pb.RegisterStreamServiceServer(grpcServer, &services.StreamServer{})
	grpcServer.Serve(listen)
}
