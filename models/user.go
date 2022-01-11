package models

import (
	"github.com/antistud/tiptoe_server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string `bson:"username,omitempty" json:"username,omitempty" binding:"required"`
	Password string `bson:"password,omitempty" json:"password,omitempty" binding:"required"`
}

func FindOneUser(user *User, id string) error {
	client, cancel, ctx := db.DbInit()

	defer client.Disconnect(ctx)
	defer cancel()

	database := client.Database("tiptoe").Collection("user")
	err := database.FindOne(ctx,
		bson.D{{"username", id}},
		options.FindOne().SetProjection(bson.M{"_id": 0, "password": 0})).Decode(user)
	if err != nil {
		return err
	}
	return nil
}
