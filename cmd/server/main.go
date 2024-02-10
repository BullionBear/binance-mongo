package main

import (
	"flag"
	"net"

	pb "github.com/BullionBear/binance-mongo/generated/proto/depth" // Adjust the import path
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDepthEventServiceServer
}

func (s *server) StreamDepthEvent(stream pb.DepthEventService_StreamDepthEventServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			glog.Infof("Finished receiving depth events: %v", err)
			break // Exit loop if stream is closed by client
		}
		glog.Infof("Received event: %v", in)
	}

	// Send a single response after all messages are received
	return stream.SendAndClose(&pb.StreamDepthEventResponse{Message: "Events Received"})
}

func main() {
	flag.Parse()       // Important: glog requires flag.Parse() to be called
	defer glog.Flush() // Ensure all logs are flushed before program exits

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDepthEventServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
