package models

import (
	"fmt"

	"github.com/antistud/tiptoe_server/db"
)

type AuthUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(user *User) error {
	client, cancel, ctx := db.DbInit()

	defer client.Disconnect(ctx)
	defer cancel()

	database := client.Database("tiptoe").Collection("user")
	res, err := database.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
