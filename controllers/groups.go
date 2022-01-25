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
	var users []models.User
	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		return
	}
	group.CreateDate = time.Now().Unix()
	createdBy := c.GetString("userId")
	group.CreatedBy = fmt.Sprintf("%v", createdBy)
	group.IsActive = true

	err := models.BulkFindUserByUsername(&users, r.Users, true)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	for _, v := range users {
		group.Users = append(group.Users, v.ID.Hex())
	}
	id, err := models.CreateMessageGroup(&group)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "groupId": id.ID.Hex()})
}
