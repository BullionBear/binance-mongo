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
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	PrevClosePrice     string `json:"x"`
	LastPrice          string `json:"c"`
	CloseQty           string `json:"Q"`
	BidPrice           string `json:"b"`
	BidQty             string `json:"B"`
	AskPrice           string `json:"a"`
	AskQty             string `json:"A"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	Count              int64  `json:"n"`
}

type WsAllMarketsStatEvent []*WsMarketStatEvent
