package routes

import (
	"server/server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Users(router *gin.Engine) {
api := router.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.POST("/users", controllers.CreateUser)
	}
	

}