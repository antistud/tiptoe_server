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

	database := db.Client.Database("tiptoe").Collection("user")
	res, err := database.InsertOne(db.Ctx, user)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
