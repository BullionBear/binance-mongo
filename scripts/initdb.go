package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/BullionBear/binance-mongo/env"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func collectionExists(ctx context.Context, db *mongo.Database, name string) bool {
	collections, err := db.ListCollectionNames(ctx, bson.M{"name": name})
	if err != nil {
		glog.Fatal(err)
	}
	return len(collections) > 0
}

func main() {
	flag.Parse() // Parse flags for glog
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.MongoURL))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("binance")

	// Create time series collection if it doesn't exist
	if !collectionExists(ctx, db, "wsDepthEvents") {
		tsOptions := options.TimeSeries().SetTimeField("E").SetMetaField("s").SetGranularity("seconds")
		collOptions := options.CreateCollection().SetTimeSeriesOptions(tsOptions)

		if err := db.CreateCollection(ctx, "wsDepthEvents", collOptions); err != nil {
			glog.Fatal(err)
		}
		glog.Infoln("Setup collection wsDepthEvents successfully.")
	}

	// Assuming index idempotency; MongoDB handles this
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "s", Value: 1},
		},
		Options: options.Index().SetUnique(false),
	}
	if _, err := db.Collection("wsDepthEvents").Indexes().CreateOne(ctx, indexModel); err != nil {
		glog.Fatal(err)
	}

	// Create rstDepthResponses collection if it doesn't exist
	if !collectionExists(ctx, db, "rstDepthResponses") {
		collOptions := options.CreateCollection()

		if err := db.CreateCollection(ctx, "rstDepthResponses", collOptions); err != nil {
			glog.Fatal(err)
		}
	}

	// Create additional indexes for rstDepthResponses
	indexModel = mongo.IndexModel{
		Keys: bson.D{
			{Key: "lastUpdateId", Value: 1},
			{Key: "symbol", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}
	if _, err := db.Collection("rstDepthResponses").Indexes().CreateOne(ctx, indexModel); err != nil {
		glog.Fatal(err)
	}

	glog.Infoln("MongoDB setup completed successfully.")
}
