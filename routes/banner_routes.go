package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBannerRoutes(router *gin.RouterGroup) {
	open := router.Group("/0/banners")
	open.GET("/all", controller.GetAllBanners)
	open.GET("/single/:id", controller.GetBanner)

	// protected routes
	protected := router.Group("/banners")
	protected.POST("/create", controller.CreateBanner)
	protected.PUT("/update/:id", controller.UpdateBanner)
	protected.DELETE("/delete/:id", controller.DeleteBanner)
}
