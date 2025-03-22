package routes

import (
	"gin-app/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("/all", controller.GetAllUsers)
		users.GET("/single/:id", controller.GetUser)
		users.POST("/create", controller.CreateUser)
		users.PUT("/update/:id", controller.UpdateUser)
		users.DELETE("/delele/:id", controller.DeleteUser)
	}
}
