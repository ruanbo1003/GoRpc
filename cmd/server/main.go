package main

import (
	"fmt"
	"google.golang.org/grpc"
	services "gorpc/cmd/server/services"
	pb "gorpc/libs/rpc/src"
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

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &services.HelloServer{})
	pb.RegisterStreamServiceServer(grpcServer, &services.StreamServer{})
	grpcServer.Serve(listen)
}
