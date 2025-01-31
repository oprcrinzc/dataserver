package db

import (
	"context"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var connString = "mongodb://localhost:27017/"
var connString string

type Config struct {
	Ip       string `toml:"ip"`
	Port     string `toml:"port"`
	DbString string `toml:"db_string"`
}

func New() *mongo.Client {
	var config Config
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Info("config file error: " + err.Error())
	}
	if config.DbString == "" {
		log.Info("Db String error")
		os.Exit(3)
	}

	connString = config.DbString

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
