package models

import (
	"context"
	"time"

	"github.com/antistud/tiptoe_server/db"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	UserId   string `bson:"userId" json:"userId"`
	IsActive bool   `bson:"isActive" json:"isActive"`
	Created  int64  `bson:"created" json:"created"`
	Token    string `bson:"token" json:"token"`
	Expires  int64  `bson:"expires" json:"expires"`
}

func CreateUser(user *User) (string, error) {

	database := db.Client.Database("tiptoe").Collection("user")
	res, err := database.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func CreateSession(username string) (string, error) {
	var session Session
	var user User
	err := FindUserByUsername(&user, username, false)
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}
	session.Token = uuid.NewString()
	session.UserId = user.ID.Hex()
	session.Created = time.Now().Unix()
	session.Expires = session.Created + 1800
	session.IsActive = true
	database := db.Client.Database("tiptoe").Collection("session")
	_, err = database.InsertOne(context.TODO(), session)
	if err != nil {
		return "", err
	}
	return session.Token, nil
}

func InvalidateSessions(userid primitive.ObjectID) error {
	// Invalidate all sessions for provided userid
	database := db.Client.Database("tiptoe").Collection("session")
	println(userid.Hex())
	res, err := database.UpdateMany(context.TODO(), bson.M{"userid": userid.Hex()}, bson.D{{"$set", bson.D{{"isactive", false}}}})
	if err != nil {
		return err
	}
	println(res.MatchedCount)
	return nil
}
