package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/antistud/tiptoe_server/db"
	"github.com/antistud/tiptoe_server/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty" binding:"required"`
	Password string             `bson:"password,omitempty" json:"password,omitempty" binding:"required"`
	Created  int64              `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func FindUserByUsername(user *User, id string, omitPassword bool) error {

	database := db.Client.Database("tiptoe").Collection("user")
	err := database.FindOne(context.TODO(),
		bson.D{{Key: "username", Value: id}},
		options.FindOne().SetProjection(bson.M{"password": util.Btoi(!omitPassword)})).Decode(user)
	if err != nil {
		return err
	}
	return nil
}

func BulkFindUserByUsername(users *[]User, usernames []string, omitPassword bool) error {
	database := db.Client.Database("tiptoe").Collection("user")
	var filter bson.A
	for _, v := range usernames {
		filter = append(filter, strings.ToLower(v))
	}
	fmt.Println(filter)
	// { field: { $in: [<value1>, <value2>, ... <valueN> ] } }
	// err := c.Find(bson.M{"friends": bson.M{"$in": arr}}).All(&users)
	c, err := database.Find(
		context.TODO(),
		bson.M{"username": bson.M{"$in": filter}},
		options.Find().SetProjection(bson.M{"password": util.Btoi(!omitPassword)}))
	if err != nil {
		return err
	}
	c.All(context.TODO(), users)
	return nil
}

func FindUserById(user *User, id string) error {
	database := db.Client.Database("tiptoe").Collection("user")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = database.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(user)
	if err != nil {
		return err
	}
	return nil
}
