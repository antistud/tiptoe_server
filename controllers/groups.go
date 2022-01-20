package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {
	var r models.GroupCreateRq
	var group models.Group

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		return
	}
	group.CreateDate = time.Now().Unix()
	createdBy := c.GetString("userId")
	group.CreatedBy = fmt.Sprintf("%v", createdBy)
	group.IsActive = true
	group.Usernames = r.Usernames
	id, err := models.CreateMessageGroup(&group)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "groupId": id.ID.Hex()})
}
