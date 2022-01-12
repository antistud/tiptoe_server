package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	var dbUser models.User
	c.BindJSON(&user)
	err := models.FindUserByUsername(&dbUser, user.Username, false)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": "invalid username or password"})
		return
	}
	passwordOk := checkPasswordHash(user.Password, dbUser.Password)
	if !passwordOk {
		fmt.Println("incorrect password")
		c.JSON(http.StatusOK, gin.H{"error": "invalid username or password"})
		return
	}
	tkn, err := models.CreateSession(user.Username)
	if err != nil {
		fmt.Println("Error creating session")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "authentication successful", "token": tkn})
}

func Logout(c *gin.Context) {
	var user models.User
	var dbUser models.User
	c.BindJSON(&user)
	err := models.FindUserByUsername(&dbUser, user.Username, true)
	if err != nil {
		fmt.Println("unable to find user")
		println(err.Error())
		c.JSON(404, gin.H{"error": "Unable to find user"})
		return
	}
	err = models.InvalidateUserSessions(dbUser.ID.Hex())
	if err != nil {
		fmt.Println("unable to invalidate sessions for the user")
		println(err.Error())
		c.JSON(404, gin.H{"error": "Unable to logout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing fields"})
		return
	}
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating User"})
		return
	}
	user.Password = hashedPass
	user.Created = time.Now().Unix()
	id, err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating User"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "userId": id})

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
