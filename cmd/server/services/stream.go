package services

import (
	pb "gorpc/libs/rpc/src"
	"io"
	"log"
)

type StreamServer struct {
}

func (s *StreamServer) List(req *pb.StreamInfoRequest, stream pb.StreamService_ListServer) error {
	for i := 0; i < 3; i++ {
		err := stream.Send(
			&pb.StreamInfoResponse{
				Item: &pb.StreamInfoItem{
					Name:  req.Item.Name,
					Value: req.Item.Value + int32(i),
				},
			})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamServer) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			item := &pb.StreamInfoItem{
				Name:  "gRPC Stream Server: Record",
				Value: 200,
			}
			return stream.SendAndClose(&pb.StreamInfoResponse{Item: item})
		}

		if err != nil {
			log.Println("StreamServer:Record, recv err:", err)
			return err
		}

		log.Printf("StreamServer:Record recv name:%s, value:%d", r.Item.Name, r.Item.Value)

	}
}

func (s *StreamServer) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		item := &pb.StreamInfoItem{
			Name:  "gRPC Stream Client: Route",
			Value: int32(n),
		}
		err := stream.Send(&pb.StreamInfoResponse{Item: item})
		if err != nil {
			log.Println("StreamServer:Route, Send err:", err)
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println("StreamServer:Route, Recv err:", err)
			return nil
		}

		n++

		log.Printf("StreamServer recv name:%s, value:%d", req.Item.Name, req.Item.Value)
	}
}
