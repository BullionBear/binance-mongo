package main

import (
	"context"
	"flag"
	"net"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/depth"
	"github.com/BullionBear/binance-mongo/utils"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	mongoURL = flag.String("mongoURL", "mongodb://localhost:27017", "MongoDB URL")
)

type server struct {
	pb.UnimplementedDepthEventServiceServer
	db *mongo.Database
}

func (s *server) StreamDepthEvent(stream pb.DepthEventService_StreamDepthEventServer) error {
	collection := s.db.Collection("wsDepthEvents")
	var buffer []interface{}
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			glog.Info("Insert documents: %v", len(buffer))
			if len(buffer) > 0 {
				_, insertErr := collection.InsertMany(context.Background(), buffer)
				if insertErr != nil {
					glog.Errorf("Failed to insert depth events into MongoDB: %v", insertErr)
				}
				// Clear the buffer after insertion
				buffer = nil
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			glog.Info("Stream closed by client")
			return nil
		default:
			in, err := stream.Recv()
			if err != nil {
				glog.Infof("Finished receiving depth events: %v", err)
				return nil // Exit loop if stream is closed by client
			}
			doc := utils.GrpcToMongoEvent(in)
			glog.Infof("Received event: %v", doc)
			buffer = append(buffer, doc)
		}
	}
}

func main() {
	flag.Parse()       // Important: glog requires flag.Parse() to be called
	defer glog.Flush() // Ensure all logs are flushed before program exits

	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI(*mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		glog.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Get a handle for your database
	db := client.Database("binance")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDepthEventServiceServer(s, &server{db: db})
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
