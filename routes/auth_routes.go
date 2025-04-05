package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.GET("/all", controller.GetAllBanners)
		auth.GET("/single/:id", controller.GetBanner)
		auth.POST("/create", controller.CreateBanner)
		auth.PUT("/update/:id", controller.UpdateBanner)
		auth.DELETE("/delete/:id", controller.DeleteBanner)
	}
}
