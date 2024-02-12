package main

import (
	"context"
	"log"
	"time"

	"github.com/BullionBear/binance-mongo/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.MongoURL))
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

	log.Println("Setup collection wsDepthEvents successfully.")

	// Create additional indexes
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"s": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if _, err := db.Collection("wsDepthEvents").Indexes().CreateOne(ctx, indexModel); err != nil {
		log.Fatal(err)
	}

	collOptions = options.CreateCollection()

	if err := db.CreateCollection(ctx, "rstDepthResponses", collOptions); err != nil {
		log.Fatal(err)
	}

	// Create additional indexes
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"lastUpdateId": 1,
			"symbol":       1,
		},
		Options: options.Index().SetUnique(true),
	}
	if _, err := db.Collection("rstDepthResponses").Indexes().CreateOne(ctx, indexModel); err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB setup completed successfully.")
}
