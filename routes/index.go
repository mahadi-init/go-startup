package routes

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	RegisterUserRoutes(router)
}
