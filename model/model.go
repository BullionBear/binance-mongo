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
