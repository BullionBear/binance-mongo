package utils

import (
	"time"

	pb "github.com/BullionBear/binance-mongo/generated/proto/wsdepth"
	"github.com/BullionBear/binance-mongo/model"
	"github.com/adshao/go-binance/v2"
)

func BinanceWsDepthToGrpcEvent(event *binance.WsDepthEvent) *pb.WsDepthEvent {
	bids := make([]*pb.Bid, len(event.Bids))
	for i, bid := range event.Bids {
		bids[i] = &pb.Bid{Price: bid.Price, Quantity: bid.Quantity}
	}
	asks := make([]*pb.Ask, len(event.Asks))
	for i, ask := range event.Asks {
		asks[i] = &pb.Ask{Price: ask.Price, Quantity: ask.Quantity}
	}
	return &pb.WsDepthEvent{
		Event:         event.Event,
		Time:          event.Time,
		Symbol:        event.Symbol,
		LastUpdateID:  event.LastUpdateID,
		FirstUpdateID: event.FirstUpdateID,
		Bids:          bids,
		Asks:          asks,
	}
}

func GrpcWsDepthToMongoEvent(event *pb.WsDepthEvent) *model.WsDepthEvent {
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
