package utils

import (
	pb "github.com/BullionBear/binance-mongo/generated/proto/wsmkstats"
	"github.com/adshao/go-binance/v2"
)

func BinanceWsWsAllMarketsStatToGrpcEvent(event *binance.WsAllMarketsStatEvent) *pb.WsAllMarketsStatEvent {
	var pbEvents []*pb.WsMarketStatEvent
	for _, e := range *event {
		pbEvent := &pb.WsMarketStatEvent{
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
		pbEvents = append(pbEvents, pbEvent)
	}
	return &pb.WsAllMarketsStatEvent{Events: pbEvents}
}
