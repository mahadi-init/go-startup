package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/0/auth")

	auth.POST("/signup", controller.Signup)
	auth.POST("/login", controller.Login)
}
