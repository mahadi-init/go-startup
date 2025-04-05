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

	// auth routes (unprotected)
	auth := router.Group("/auth")
	{
		RegisterAuthRoutes(auth)
	}

	// api routes (protected)
	api := router.Group("/api")
	{
		RegisterUserRoutes(api)
		RegisterBannerRoutes(api)
	}
}
