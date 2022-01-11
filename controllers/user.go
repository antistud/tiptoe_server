package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Query("username")
	if id == "" {
		fmt.Println("Incorrect url params")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}
	fmt.Println("username param")
	fmt.Println(id)
	err := models.FindOneUser(&user, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, errors.New("can't find user"))
		return
	}
	c.JSON(http.StatusOK, user)
}
