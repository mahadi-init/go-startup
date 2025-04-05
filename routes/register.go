package routes

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	root := router.Group("/api")
	RegisterAuthRoutes(root)
	RegisterUserRoutes(root)
	RegisterBannerRoutes(root)
}
