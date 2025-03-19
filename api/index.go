package api

import (
	"gin-app/controllers"
	"gin-app/middleware"
	"gin-app/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter creates and configures a Gin router
func SetupRouter() *gin.Engine {
	// Set Gin to release mode in production
	// gin.SetMode(gin.ReleaseMode) // Uncomment this for production

	// Create a new Gin engine
	router := gin.New()
	router.Use(gin.Recovery())

	// Apply global middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Initialize controllers
	userController := controllers.NewUserController()

	// Setup routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go API with Gin on Vercel!",
		})
	})

	// Register user routes
	userRoutes := router.Group("/api")
	routes.RegisterUserRoutes(userRoutes, userController)
	return router
}

// Handler function for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the router
	router := SetupRouter()

	// Serve the request
	router.ServeHTTP(w, r)
}
