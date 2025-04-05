package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	open := router.Group("/0/auth")

	open.POST("/signup", controller.Signup)
	open.POST("/login", controller.Login)
}
