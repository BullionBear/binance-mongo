package utils

import (
	pb "github.com/BullionBear/binance-mongo/generated/proto/wspdepth"
	"github.com/adshao/go-binance/v2"
)

func BinanceWsPartialDepthToGrpcEvent(event *binance.WsPartialDepthEvent) *pb.WsPartialDepthEvent {
	bids := make([]*pb.Bid, len(event.Bids))
	for i, bid := range event.Bids {
		bids[i] = &pb.Bid{Price: bid.Price, Quantity: bid.Quantity}
	}
	asks := make([]*pb.Ask, len(event.Asks))
	for i, ask := range event.Asks {
		asks[i] = &pb.Ask{Price: ask.Price, Quantity: ask.Quantity}
	}
	return &pb.WsPartialDepthEvent{
		Symbol:       event.Symbol,
		LastUpdateID: event.LastUpdateID,
		Bids:         bids,
		Asks:         asks,
	}
}
