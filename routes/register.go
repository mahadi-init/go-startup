package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go API with Gin on Vercel!",
		})
	})

	root := router.Group("/api")
	RegisterAuthRoutes(root)
	RegisterUserRoutes(root)
	RegisterBannerRoutes(root)
}
