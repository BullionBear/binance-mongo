package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27016"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("binance")

	// Create time series collection
	tsOptions := options.TimeSeries().SetTimeField("E").SetMetaField("s").SetGranularity("seconds")
	collOptions := options.CreateCollection().SetTimeSeriesOptions(tsOptions)

	if err := db.CreateCollection(ctx, "wsDepthEvents", collOptions); err != nil {
		log.Fatal(err)
	}

	// Create additional indexes
	// indexModel := mongo.IndexModel{
	// 	Keys:    bson.D{{"Bids.Price", 1}}, // Example index
	// 	Options: options.Index().SetUnique(false),
	// }
	// if _, err := db.Collection("depthEvents").Indexes().CreateOne(ctx, indexModel); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("MongoDB setup completed successfully.")
}
