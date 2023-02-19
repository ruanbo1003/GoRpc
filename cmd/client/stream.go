package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "gorpc/libs/rpc/src"
	"io"
	"log"
)

type StreamClient struct{}

func (s *StreamClient) ReqLists(conn *grpc.ClientConn, ctx context.Context) {
	req := &pb.StreamInfoRequest{
		Item: &pb.StreamInfoItem{
			Name:  "stream-client request: List",
			Value: 10,
		}}
	client := pb.NewStreamServiceClient(conn)
	stream, err := client.List(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("list rsp, name:%s, value:%d", rsp.Item.Name, rsp.Item.Value)
	}
}

func (s *StreamClient) ReqRecord(conn *grpc.ClientConn, ctx context.Context) {
	client := pb.NewStreamServiceClient(conn)
	stream, err := client.Record(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &pb.StreamInfoRequest{
		Item: &pb.StreamInfoItem{
			Name:  "stream-client request: Record",
			Value: 20,
		}}
	for i := 0; i < 3; i++ {
		err := stream.Send(req)
		if err != nil {
			fmt.Println("stream-client-Record:Send", err)
			return
		}
	}

	rsp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("stream-client-Record:CloseAndRecd", err)
		return
	}

	log.Printf("rsp, name:%s, value:%d", rsp.Item.Name, rsp.Item.Value)
}

func (s *StreamClient) ReqRoute(conn *grpc.ClientConn, ctx context.Context) {
	client := pb.NewStreamServiceClient(conn)
	stream, err := client.Route(ctx)
	if err != nil {
		return
	}

	req := &pb.StreamInfoRequest{
		Item: &pb.StreamInfoItem{
			Name:  "stream-client request: Route",
			Value: 20,
		}}

	for i := 0; i < 3; i++ {
		err = stream.Send(req)
		if err != nil {
			log.Println("stream-client:Route, send err:", err)
			return
		}

		rsp, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("stream-client:Route, recv err:", err)
		}

		log.Printf("stream-client:Route, recv name:%s, value:%d", rsp.Item.Name, rsp.Item.Value)
	}
}
