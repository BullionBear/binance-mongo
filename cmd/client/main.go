package main

import (
	"context"
	"flag"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/depth" // Adjust the import path
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse() // Parse flags for glog
	defer glog.Flush()

	// Establish a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		glog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDepthEventServiceClient(conn)

	// Create a stream.
	stream, err := c.StreamDepthEvent(context.Background())
	if err != nil {
		glog.Fatalf("could not stream: %v", err)
	}

	// Send messages.
	for i := 0; i < 100; i++ {
		event := &pb.WsDepthEvent{
			Event:         "depthUpdate",
			Time:          time.Now().Unix(),
			Symbol:        "BTCUSDT",
			LastUpdateID:  int64(i),
			FirstUpdateID: int64(i - 1),
			Bids:          []*pb.Bid{{Price: "50000.0", Quantity: "1.0"}},
			Asks:          []*pb.Ask{{Price: "50001.0", Quantity: "1.0"}},
		}
		if err := stream.Send(event); err != nil {
			glog.Fatalf("Failed to send a depth event: %v", err)
		}
		glog.Infof("Sent depth event: %v", event)
	}

	// Close the stream and receive the server's response.
	response, err := stream.CloseAndRecv()
	if err != nil {
		glog.Fatalf("Failed to receive response: %v", err)
	}
	glog.Infof("Server response: %v", response)
}
