package controllers

import (
	"fmt"
	"net/http"

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
