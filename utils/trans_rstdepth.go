package utils

import (
	pb "github.com/BullionBear/binance-mongo/generated/proto/rst_depth"
	"github.com/BullionBear/binance-mongo/model"
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

func GrpcDepthToMongoEvent(event *pb.DepthResponse) *model.DepthResponse {
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
	return &model.DepthResponse{
		LastUpdateID: event.LastUpdateID,
		Symbol:       event.Symbol,
		Bids:         bids,
		Asks:         asks,
	}
}
