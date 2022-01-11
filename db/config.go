package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func DbInit() (context.Context, context.CancelFunc, error) {

	/*
	   Connect to my cluster
	*/
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("TIPTOE_DB_URI")))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}
	return ctx, cancel, nil

}
