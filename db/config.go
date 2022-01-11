package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var Ctx context.Context
var CtxCancel context.CancelFunc

func DbInit() (*mongo.Client, context.CancelFunc, context.Context) {

	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("TIPTOE_DB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, cancel, ctx

}
