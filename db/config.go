package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Ctx context.Context
var CtxCancel context.CancelFunc

func DbInit() error {

	/*
	   Connect to my cluster
	*/
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("TIPTOE_DB_URI")))
	if err != nil {
		return err
	}
	Ctx, CtxCancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(Ctx)
	if err != nil {
		return err
	}
	return nil

}
