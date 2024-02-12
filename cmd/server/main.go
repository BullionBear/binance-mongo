package main

import (
	"context"
	"flag"
	"net"

	"github.com/BullionBear/binance-mongo/env"
	"github.com/BullionBear/binance-mongo/services/rstdepth"
	"github.com/BullionBear/binance-mongo/services/wsdepth"
	"github.com/BullionBear/binance-mongo/utils"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	mongoURL = flag.String("mongoURL", env.MongoURL, "MongoDB URL")
)

func main() {
	utils.PrintEnv("Server")
	flag.Parse()
	glog.Infoln("Host on: ", "tcp:50051")
	glog.Infoln("MongoDB URL: ", *mongoURL)
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
