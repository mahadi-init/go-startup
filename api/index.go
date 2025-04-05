package api

import (
	"gin-app/db"
	"gin-app/middleware"
	"gin-app/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter creates and configures a Gin router
func SetupRouter() *gin.Engine {
	// Set Gin to release mode in production
	gin.SetMode(gin.ReleaseMode) // Uncomment this for production

	// Initialize DB connection
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err)
	}

	// Create a new Gin engine
	router := gin.New()
	router.Use(gin.Recovery())

	// home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go API with Gin on Vercel!",
		})
	})

	// Apply global middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.JWTAuthMiddleware())

	// Register routes
	routes.Register(router)
	return router
}

// Handler function for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the router
	router := SetupRouter()

	// Serve the request
	router.ServeHTTP(w, r)
}
