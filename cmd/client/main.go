package main

import (
	"context"
	"flag"

	pb "github.com/BullionBear/binance-mongo/generated/proto/depth"
	"github.com/BullionBear/binance-mongo/utils"
	"github.com/adshao/go-binance/v2"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func main() {
	symbol := flag.String("symbol", "BTCUSDT", "Trading symbol")
	grpcServerAddr := flag.String("grpc-server", "localhost:50051", "gRPC server address")

	flag.Parse() // Parse flags
	defer glog.Flush()

	// Establish a connection to the server.
	conn, err := grpc.Dial(*grpcServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		glog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDepthEventServiceClient(conn)

	// Create a stream to the gRPC server.
	stream, err := client.StreamDepthEvent(context.Background())
	if err != nil {
		glog.Fatalf("could not create stream: %v", err)
	}

	// Connect to Binance WebSocket for depth events.
	doneC, _, err := binance.WsDepthServe(*symbol, func(event *binance.WsDepthEvent) {
		grpcEvent := utils.ConvertToGrpcEvent(event)
		if err := stream.Send(grpcEvent); err != nil {
			glog.Errorf("Failed to send depth event to gRPC server: %v", err)
		} else {
			glog.Infof("Sent depth event to gRPC server: %v", grpcEvent)
		}
	}, func(err error) {
		glog.Errorf("WebSocket Error: %v", err)
	})

	if err != nil {
		glog.Fatalf("Failed to connect to Binance WebSocket: %v", err)
	}

	<-doneC // Keep the connection alive.

	// Close the stream and receive the server's response after doneC is closed (indicating WebSocket closure).
	if response, err := stream.CloseAndRecv(); err != nil {
		glog.Fatalf("Failed to receive closing response: %v", err)
	} else {
		glog.Infof("Server closing response: %v", response)
	}
}
