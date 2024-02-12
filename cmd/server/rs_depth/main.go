package main

import (
	"context"
	"flag"
	"net"
	"sync"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/rstdepth"
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
	pb.UnimplementedDepthResponseServiceServer
	db *mongo.Database
	mu sync.Mutex // Mutex to protect the buffer
}

func (s *server) flushBuffer(buffer *[]interface{}, collection *mongo.Collection) {
	s.mu.Lock() // Ensure exclusive access to the buffer
	defer s.mu.Unlock()
	n_doc := len(*buffer)
	glog.Infof("Number of documents are inserted: %v", n_doc)
	if n_doc > 0 {
		_, err := collection.InsertMany(context.Background(), *buffer)
		if err != nil {
			glog.Errorf("Failed to insert depth events into MongoDB: %v", err)
		}
		*buffer = (*buffer)[:0] // Efficiently clear the buffer while retaining allocated memory
	}
}

func (s *server) StreamDepthResponse(stream pb.DepthResponseService_StreamDepthResponseServer) error {
	collection := s.db.Collection("rstDepthResponses")
	buffer := make([]interface{}, 0, 1024) // Preallocate buffer with estimated capacity
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.flushBuffer(&buffer, collection)
		case <-stream.Context().Done():
			glog.Info("Stream closed by client")
			s.flushBuffer(&buffer, collection)
			return nil
		default:
			in, err := stream.Recv()
			if err != nil {
				glog.Infof("Finished receiving depth events: %v", err)
				s.flushBuffer(&buffer, collection) // Ensure buffer is flushed before exiting
				return nil
			}
			doc := utils.GrpcDepthToMongoEvent(in)
			// glog.Infof("Received event: %v", doc)

			s.mu.Lock()
			buffer = append(buffer, doc)
			if len(buffer) >= 10 {
				s.mu.Unlock() // Unlock before flushing to avoid deadlock
				s.flushBuffer(&buffer, collection)
			} else {
				s.mu.Unlock()
			}
		}
	}
}

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
	pb.RegisterDepthResponseServiceServer(s, &server{db: db})
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("Failed to serve: %v", err)
	}
}
