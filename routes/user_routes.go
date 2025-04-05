package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	protected := router.Group("/users")

	protected.GET("/all", controller.GetAllUsers)
	protected.GET("/single/:id", controller.GetUser)
	protected.POST("/create", controller.CreateUser)
	protected.PUT("/update/:id", controller.UpdateUser)
	protected.DELETE("/delete/:id", controller.DeleteUser)
}
