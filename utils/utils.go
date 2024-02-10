package utils

import (
	pb "github.com/BullionBear/binance-mongo/generated/proto/depth"
	"github.com/BullionBear/binance-mongo/model"
	"github.com/adshao/go-binance/v2"
)

func BinanceToGrpcEvent(event *binance.WsDepthEvent) *pb.WsDepthEvent {
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

func GrpcToMongoEvent(event *pb.WsDepthEvent) *model.WsDepthEvent {
	// Allocate slices for bids and asks with the correct length.
	bids := make([]model.Bid, len(event.Bids))
	for i, bid := range event.Bids {
		// Assuming bid.Price and bid.Quantity are the correct fields and types.
		// Directly assign to the slice without using pointers since model.Bid is a type alias for PriceLevel, not a pointer.
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
		Time:          event.Time,
		Symbol:        event.Symbol,
		LastUpdateID:  event.LastUpdateID,
		FirstUpdateID: event.FirstUpdateID,
		Bids:          bids,
		Asks:          asks,
	}
}
