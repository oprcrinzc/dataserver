package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connString = "mongodb://localhost:27017/"

func New() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return client
}
