package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	var r models.MessageCreateRq
	var message models.Message
	r.ToGroup = "61ea37c642ab6e44ed8572d0"
	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		return
	}
	message.CreateDate = time.Now().Unix()
	sender := c.GetString("userId")
	message.Sender = fmt.Sprintf("%v", sender)
	message.ToGroup = r.ToGroup
	// TODO: we gotta upload the data in the request to S3 and then get the destination url
	// using dummy for now
	message.VideoUrl = "www.youtube.com"
	id, err := models.CreateMessage(&message)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "messageId": id.ID.Hex()})
}
