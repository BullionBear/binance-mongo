package model

import "time"

type PriceLevel struct {
	Price    string
	Quantity string
}

// Ask is a type alias for PriceLevel.
type Ask = PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = PriceLevel

// WsDepthEvent define websocket depth event
type WsDepthEvent struct {
	Event         string    `json:"e" bson:"e"`
	Time          time.Time `json:"E" bson:"E"` // Apply for MongoDB time series
	Symbol        string    `json:"s" bson:"s"`
	LastUpdateID  int64     `json:"u" bson:"u"`
	FirstUpdateID int64     `json:"U" bson:"U"`
	Bids          []Bid     `json:"b" bson:"b"`
	Asks          []Ask     `json:"a" bson:"a"`
}

type DepthResponse struct {
	LastUpdateID int64  `json:"lastUpdateId" bson:"lastUpdateId"`
	Symbol       string `json:"symbol" bson:"symbol"`
	Bids         []Bid  `json:"bids" bson:"bids"`
	Asks         []Ask  `json:"asks" bson:"asks"`
}

type WsPartialDepthEvent = DepthResponse

type WsMarketStatEvent struct {
	Event              string    `json:"e" bson:"e"`
	Time               time.Time `json:"E" bson:"E"`
	Symbol             string    `json:"s" bson:"s"`
	PriceChange        string    `json:"p" bson:"p"`
	PriceChangePercent string    `json:"P" bson:"P"`
	WeightedAvgPrice   string    `json:"w" bson:"w"`
	PrevClosePrice     string    `json:"x" bson:"x"`
	LastPrice          string    `json:"c" bson:"c"`
	CloseQty           string    `json:"Q" bson:"Q"`
	BidPrice           string    `json:"b" bson:"b"`
	BidQty             string    `json:"B" bson:"B"`
	AskPrice           string    `json:"a" bson:"a"`
	AskQty             string    `json:"A" bson:"A"`
	OpenPrice          string    `json:"o" bson:"o"`
	HighPrice          string    `json:"h" bson:"h"`
	LowPrice           string    `json:"l" bson:"l"`
	BaseVolume         string    `json:"v" bson:"v"`
	QuoteVolume        string    `json:"q" bson:"q"`
	OpenTime           time.Time `json:"O" bson:"O"`
	CloseTime          time.Time `json:"C" bson:"C"`
	FirstID            int64     `json:"F" bson:"F"`
	LastID             int64     `json:"L" bson:"L"`
	Count              int64     `json:"n" bson:"n"`
}

type WsAllMarketsStatEvent []*WsMarketStatEvent
