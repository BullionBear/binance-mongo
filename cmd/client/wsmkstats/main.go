package main

import (
	"context"
	"flag"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/wsmkstats"
	"github.com/BullionBear/binance-mongo/utils"
	"github.com/adshao/go-binance/v2"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcServerAddr := flag.String("grpc-server", "localhost:50051", "gRPC server address")

	flag.Parse() // Parse flags
	utils.PrintEnv("Client WS Markets Stat")
	glog.Infoln("Connect to: ", *grpcServerAddr)
	defer glog.Flush()

	// Establish a connection to the server.
	conn, err := grpc.Dial(*grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		glog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAllMarketsStatEventServiceClient(conn)

	// Create a stream to the gRPC server.
	stream, err := client.StreamAllMarketsStatEvent(context.Background())
	if err != nil {
		glog.Fatalf("could not create stream: %v", err)
	}

	utils.EchoClock(30 * time.Second)

	// Connect to Binance WebSocket for depth events.
	doneC, _, err := binance.WsAllMarketsStatServe(func(event binance.WsAllMarketsStatEvent) {
		grpcEvent := utils.BinanceWsWsAllMarketsStatToGrpcEvent(&event)
		utils.IncrementCounter()
		if err := stream.Send(grpcEvent); err != nil {
			glog.Errorf("Failed to send markets stat event to gRPC server: %v", err)
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
