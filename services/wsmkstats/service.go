package wsmkstats

import (
	"context"
	"sync"
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/wsmkstats"
	"github.com/BullionBear/binance-mongo/model"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAllMarketsStatEventServiceServer
	Db *mongo.Database
	mu sync.Mutex // Mutex to protect the buffer
}

func (s *Server) StreamAllMarketsStatEvent(stream pb.AllMarketsStatEventService_StreamAllMarketsStatEventServer) error {
	collection := s.Db.Collection("wsMarketStatEvents") // Should be the only difference between wsdepth
	buffer := make([]interface{}, 0, 128)               // Preallocate buffer with estimated capacity
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
			if len(buffer) >= 64 {
				s.mu.Unlock() // Unlock before flushing to avoid deadlock
				s.flushBuffer(&buffer, collection)
			} else {
				s.mu.Unlock()
			}
		}
	}
}

func Register(grpc *grpc.Server, s *Server) {
	pb.RegisterAllMarketsStatEventServiceServer(grpc, s)
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

// toDoc converts a protobuf WsAllMarketsStatEvent to a model WsAllMarketsStatEvent.
func toDoc(event *pb.WsAllMarketsStatEvent) *model.WsAllMarketsStatEvent {
	var modelEvents model.WsAllMarketsStatEvent
	for _, e := range event.Events {
		modelEvent := &model.WsMarketStatEvent{
			Event:              e.Event,
			Time:               e.Time,
			Symbol:             e.Symbol,
			PriceChange:        e.PriceChange,
			PriceChangePercent: e.PriceChangePercent,
			WeightedAvgPrice:   e.WeightedAvgPrice,
			PrevClosePrice:     e.PrevClosePrice,
			LastPrice:          e.LastPrice,
			CloseQty:           e.CloseQty,
			BidPrice:           e.BidPrice,
			BidQty:             e.BidQty,
			AskPrice:           e.AskPrice,
			AskQty:             e.AskQty,
			OpenPrice:          e.OpenPrice,
			HighPrice:          e.HighPrice,
			LowPrice:           e.LowPrice,
			BaseVolume:         e.BaseVolume,
			QuoteVolume:        e.QuoteVolume,
			OpenTime:           e.OpenTime,
			CloseTime:          e.CloseTime,
			FirstID:            e.FirstID,
			LastID:             e.LastID,
			Count:              e.Count,
		}
		modelEvents = append(modelEvents, modelEvent)
	}
	return &modelEvents
}
