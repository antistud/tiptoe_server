package routes

import (
	"github.com/antistud/tiptoe_server/controllers"
	"github.com/antistud/tiptoe_server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("/user/")
		user.Use(middleware.AuthRequired())
		{
			user.GET("/userByName", controllers.GetUser)
		}

		auth := v1.Group("/auth/")
		{
			auth.POST("login", controllers.Login)
			auth.POST("createUser", controllers.CreateUser)
			auth.POST("logout", controllers.Logout)
		}

		connection := v1.Group("/connection/")
		connection.Use(middleware.AuthRequired())
		{
			connection.POST("/create")
		}

		message := v1.Group("/message/")
		message.Use(middleware.AuthRequired())
		{
			message.POST("/createGroup", controllers.CreateGroup)
			message.POST("/create", middleware.CheckGroup(), controllers.CreateMessage)
		}
	}
}
