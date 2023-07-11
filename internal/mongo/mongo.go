package mongodb

import (
	"context"
	"fmt"
	"log"
	"simplepatientorder/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine(config *config.Config) {
	var err error
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.DB.Mongo.Host, config.DB.Mongo.Port))
	clientOptions.SetServerSelectionTimeout(time.Second * time.Duration(config.DB.Mongo.ServerSelectionTimeoutSec))
	clientOptions.SetConnectTimeout(time.Second * time.Duration(config.DB.Mongo.ConnectTimeoutSec))
	clientOptions.SetTimeout(time.Second * time.Duration(config.DB.Mongo.OperationTimeoutSec))

	mgoCli, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect mongo:", err)
	}

	err = mgoCli.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping mongo:", err)
	}
}

func GetMgoCli(config *config.Config) *mongo.Client {
	if mgoCli == nil {
		initEngine(config)
	}

	return mgoCli
}
