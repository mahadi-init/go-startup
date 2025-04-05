package routes

import (
	controller "gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBannerRoutes(router *gin.RouterGroup) {
	banners := router.Group("/banners")
	{
		banners.GET("/all", controller.GetAllBanners)
		banners.GET("/single/:id", controller.GetBanner)
		banners.POST("/create", controller.CreateBanner)
		banners.PUT("/update/:id", controller.UpdateBanner)
		banners.DELETE("/delete/:id", controller.DeleteBanner)
	}
}
