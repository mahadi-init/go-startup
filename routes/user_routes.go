package routes

import (
	"gin-app/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(router *gin.RouterGroup, controller *controllers.UserController) {
	users := router.Group("/users")
	{
		users.GET("", controller.GetUsers)
		users.GET("/:id", controller.GetUser)
		users.POST("", controller.CreateUser)
		users.PUT("/:id", controller.UpdateUser)
		users.DELETE("/:id", controller.DeleteUser)
	}
}
