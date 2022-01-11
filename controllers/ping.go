package controllers

import (
	"net/http"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	res := models.Pong{Title: "TITLE", Description: "PONG"}
	c.JSON(http.StatusOK, res)
}
