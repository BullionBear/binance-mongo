package wsdepth

import (
	"context"
	"sync"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/wsdepth"
	"github.com/BullionBear/binance-mongo/model"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDepthEventServiceServer
	Db *mongo.Database
	mu sync.Mutex // Mutex to protect the buffer
}

func (s *Server) StreamDepthEvent(stream pb.DepthEventService_StreamDepthEventServer) error {
	collection := s.Db.Collection("wsDepthEvents")
	buffer := make([]interface{}, 0, 4096) // Preallocate buffer with estimated capacity
	ticker := time.NewTicker(10 * time.Second)
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
			doc := toDoc(in)
			// glog.Infof("Received event: %v", doc)

			s.mu.Lock()
			buffer = append(buffer, doc)
			if len(buffer) >= 1024 {
				s.mu.Unlock() // Unlock before flushing to avoid deadlock
				s.flushBuffer(&buffer, collection)
			} else {
				s.mu.Unlock()
			}
		}
	}
}

func Register(grpc *grpc.Server, s *Server) {
	pb.RegisterDepthEventServiceServer(grpc, s)
	return
}

func (s *Server) flushBuffer(buffer *[]interface{}, collection *mongo.Collection) {
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

func toDoc(event *pb.WsDepthEvent) *model.WsDepthEvent {
	bids := make([]model.Bid, len(event.Bids))
	for i, bid := range event.Bids {
		bids[i] = model.Bid{Price: bid.Price, Quantity: bid.Quantity}
	}

	asks := make([]model.Ask, len(event.Asks))
	for i, ask := range event.Asks {
		// Similar handling for asks as for bids.
		asks[i] = model.Ask{Price: ask.Price, Quantity: ask.Quantity}
	}

	// Return a pointer to the constructed WsDepthEvent, filled with the converted slices.
	return &model.WsDepthEvent{
		Event:         event.Event,
		Time:          time.Unix(0, event.Time*int64(time.Millisecond)),
		Symbol:        event.Symbol,
		LastUpdateID:  event.LastUpdateID,
		FirstUpdateID: event.FirstUpdateID,
		Bids:          bids,
		Asks:          asks,
	}
}
