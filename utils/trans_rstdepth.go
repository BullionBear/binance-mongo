package utils

import (
	pb "github.com/BullionBear/binance-mongo/generated/proto/rstdepth"
	"github.com/adshao/go-binance/v2"
)

func BinanceDepthToGrpcEvent(response *binance.DepthResponse, symbol string) *pb.DepthResponse {
	bids := make([]*pb.Bid, len(response.Bids))
	for i, bid := range response.Bids {
		bids[i] = &pb.Bid{Price: bid.Price, Quantity: bid.Quantity}
	}
	asks := make([]*pb.Ask, len(response.Asks))
	for i, ask := range response.Asks {
		asks[i] = &pb.Ask{Price: ask.Price, Quantity: ask.Quantity}
	}
	return &pb.DepthResponse{
		LastUpdateID: response.LastUpdateID,
		Symbol:       symbol,
		Bids:         bids,
		Asks:         asks,
	}
}
