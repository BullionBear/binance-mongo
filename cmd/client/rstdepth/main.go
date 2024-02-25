package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/rstdepth"
	"github.com/BullionBear/binance-mongo/utils"
	"github.com/adshao/go-binance/v2"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	symbol := flag.String("symbol", "BTCUSDT", "Trading symbol")
	grpcServerAddr := flag.String("grpc-server", "localhost:50051", "gRPC server address")

	flag.Parse() // Parse flags
	utils.PrintEnv("Client RST Depth")
	glog.Infoln("Symbol: ", symbol)
	glog.Infoln("Connect to: ", grpcServerAddr)
	defer glog.Flush()

	// Establish a connection to the server.
	conn, err := grpc.Dial(*grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		glog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gclient := pb.NewDepthResponseServiceClient(conn)

	utils.EchoClock(30 * time.Second)
	// Create a stream to the gRPC server.
	stream, err := gclient.StreamDepthResponse(context.Background())
	if err != nil {
		glog.Fatalf("failed to create stream: %v", err)
	}
	defer stream.CloseSend() // Ensure to close the stream when done.
	bclient := binance.NewClient("", "")

	// Create a ticker and a context to handle graceful shutdown.
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Create a channel to listen for interrupt signal to gracefully shut down.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Use a context to cancel operations when the application is asked to shut down.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-stopChan // Wait for interrupt signal
		glog.Info("Shutting down gracefully...")
		ticker.Stop() // Stop the ticker
		cancel()      // Cancel the context to clean up resources
	}()

	for {
		select {
		case <-ctx.Done():
			glog.Info("Application stopped.")
			return
		case <-ticker.C:
			res, err := bclient.NewDepthService().Symbol(*symbol).Limit(50).
				Do(context.Background())
			if err != nil {
				glog.Error(err)
				continue
			}
			grpcEvent := utils.BinanceDepthToGrpcEvent(res, *symbol)
			utils.IncrementCounter()
			if err := stream.Send(grpcEvent); err != nil {
				glog.Errorf("Failed to send depth event to gRPC server: %v", err)
			} else {
				glog.Infof("Sent depth event to gRPC server: %v", grpcEvent)
			}
		}
	}
}
