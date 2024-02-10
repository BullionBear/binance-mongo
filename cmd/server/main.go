package main

import (
	"context"
	"flag"
	"net"

	pb "github.com/BullionBear/binance-mongo/generated/proto/depth"
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
	collection := s.db.Collection("depth")
	for {
		in, err := stream.Recv()
		if err != nil {
			glog.Infof("Finished receiving depth events: %v", err)
			break // Exit loop if stream is closed by client
		}
		glog.Infof("Received event: %v", in)
		// Insert the event into MongoDB
		_, insertErr := collection.InsertOne(context.Background(), in)
		if insertErr != nil {
			glog.Errorf("Failed to insert depth event into MongoDB: %v", insertErr)
		}
	}
	// Send a single response after all messages are received
	return stream.SendAndClose(&pb.StreamDepthEventResponse{Message: "Events Received"})
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
