package database

import (
	"context"

	"github.com/P-Tesch/go-svelte/backend/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Database

func GetDatabase() *mongo.Database {
	if db == nil {
		setup()
	}
	return db
}

func setup() {
	uri := helpers.GetEnvVar("MONGODB_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	helpers.HandleError(err)

	var result bson.M
	err = client.Database("main").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result)
	helpers.HandleError(err)

	db = client.Database("main")
}
