package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user models.User
	username := c.Query("username")
	id := c.Query("id")
	if id == "" && username == "" {
		fmt.Println("Incorrect url params")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}
	if id == "" && username != "" {
		err := models.FindUserByUsername(&user, username, true)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	if id != "" && username == "" {
		err := models.FindUserById(&user, id)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		return
	}
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Password = hashedPass
	user.Created = time.Now().Unix()
	id, err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "userId": id})

}
