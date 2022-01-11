package controllers

import (
	"fmt"
	"net/http"

	"github.com/antistud/tiptoe_server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	var dbUser models.User
	c.BindJSON(&user)
	err := models.FindOneUser(&dbUser, user.Username)
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
	c.JSON(http.StatusOK, gin.H{"status": "authentication successful"})
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
	err = models.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating User"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Successfully Created"})

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
