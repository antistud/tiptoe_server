package middleware

import (
	"net/http"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		userId, err := models.IsSessionValid(token)
		if token == "" || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("userId", userId)
	}
}

func CheckGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: check to make sure requestor has permission to send message here
	}
}
