package main

import (
	"context"
	"flag"
	"net"

	"github.com/BullionBear/binance-mongo/services/rstdepth"
	"github.com/BullionBear/binance-mongo/services/wsdepth"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	mongoURL = flag.String("mongoURL", "mongodb://localhost:27016", "MongoDB URL")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	clientOptions := options.Client().ApplyURI(*mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		glog.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("binance")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		glog.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	wsdepth.Register(s, &wsdepth.Server{Db: db})
	rstdepth.Register(s, &rstdepth.Server{Db: db})
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("Failed to serve: %v", err)
	}
}
