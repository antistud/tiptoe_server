package models

import (
	"context"
	"fmt"

	"github.com/antistud/tiptoe_server/db"
)

type AuthUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(user *User) error {

	database := db.Client.Database("tiptoe").Collection("user")
	res, err := database.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
