package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Logger is a middleware that logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request details with colorized output
		fmt.Printf("\n\033[1;34m[%s]\033[0m \033[1;32m%s\033[0m \033[1;36m%s\033[0m \033[1;33m%d\033[0m \033[1;31m%s\033[0m\n",
			time.Now().Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)
	}
}
