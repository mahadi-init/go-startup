package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")

	users.GET("/all", controller.GetAllUsers)
	users.GET("/single/:id", controller.GetUser)
	users.POST("/create", controller.CreateUser)
	users.PUT("/update/:id", controller.UpdateUser)
	users.DELETE("/delete/:id", controller.DeleteUser)
}
