package util

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func UnaryLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC request[%s]", info.FullMethod)
	return handler(ctx, req)
}

func StreamLoggingInterceptor(src interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("gRPC stream request[%s]", info.FullMethod)
	return handler(src, ss)
}
