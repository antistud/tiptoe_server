package models

import (
	"context"
	"errors"
	"time"

	"github.com/antistud/tiptoe_server/db"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	UserId  string `bson:"userId" json:"userId"`
	Created int64  `bson:"created" json:"created"`
	Token   string `bson:"token" json:"token"`
	Expires int64  `bson:"expires" json:"expires"`
}

type LogoutRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
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
	database := db.Client.Database("tiptoe").Collection("session")
	_, err = database.InsertOne(context.TODO(), session)
	if err != nil {
		return "", err
	}
	return session.Token, nil
}

func InvalidateUserSessions(userid string) error {
	// Invalidate all sessions for provided userid
	database := db.Client.Database("tiptoe").Collection("session")
	_, err := database.UpdateMany(context.TODO(), bson.D{{"userId", userid}}, bson.D{{"$set", bson.D{{"expires", 0}}}})
	if err != nil {
		return err
	}
	return nil
}

func IsUserSessionValid(token string, userid string) error {
	database := db.Client.Database("tiptoe").Collection("session")
	var res Session
	var filter bson.D
	if userid == "" {
		filter = bson.D{{"token", token}}
	} else {
		filter = bson.D{{"token", token}, {"userId", userid}}
	}
	err := database.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return err
	}
	if time.Now().Unix() > res.Expires {
		return errors.New("Token Inactive")
	}
	return nil
}

func IsSessionValid(token string) (string, error) {
	err := IsUserSessionValid(token, "")
	if err != nil {
		return "", err
	}
	return "a", nil
}

func GetUserFromSession(token string) (string, error) {
	database := db.Client.Database("tiptoe").Collection("session")
	var user User
	filter := bson.D{{"token", token}}
	err := database.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.ID.Hex(), nil
}
