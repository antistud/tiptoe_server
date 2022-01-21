package models

import (
	"context"

	"github.com/antistud/tiptoe_server/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	CreateDate int64  `bson:"createDate" json:"createDate"`
	Sender     string `bson:"sentBy" json:"sentBy"`
	ToGroup    string `bson:"toGroup" json:"toGroup"`
	VideoUrl   string `bson:"videoUrl" json:"videoUrl"`
}

type MessageCreateRs struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

type MessageCreateRq struct {
	ToGroup   string `bson:"toGroup" json:"toGroup"`
	VideoData string `bson:"videoData" json:"videoData"`
}

func CreateMessage(r *Message) (MessageCreateRs, error) {

	database := db.Client.Database("tiptoe").Collection("message")

	res, err := database.InsertOne(context.TODO(), r)
	if err != nil {
		return MessageCreateRs{primitive.NilObjectID}, err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return MessageCreateRs{oid}, nil
}
