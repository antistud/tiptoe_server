package models

import (
	"context"

	"github.com/antistud/tiptoe_server/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	CreatedBy  string   `json:"createdBy"`
	CreateDate int64    `json:"createDate"`
	Usernames  []string `json:"usernames"`
	IsActive   bool     `json:"isActive"`
}

type GroupCreateRq struct {
	Usernames  []string `json:"usernames"`
	CreateDate uint64   `json:"createDate,omitempty"`
}

type GroupCreateRs struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

func CreateMessageGroup(r *Group) (GroupCreateRs, error) {

	database := db.Client.Database("tiptoe").Collection("groups")

	res, err := database.InsertOne(context.TODO(), r)
	if err != nil {
		return GroupCreateRs{primitive.NilObjectID}, err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return GroupCreateRs{oid}, nil
}
