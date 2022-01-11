package models

import (
	"context"

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

	database := db.Client.Database("tiptoe").Collection("user")
	err := database.FindOne(context.TODO(),
		bson.D{{"username", id}},
		options.FindOne().SetProjection(bson.M{"_id": 0, "password": 0})).Decode(user)
	if err != nil {
		return err
	}
	return nil
}
