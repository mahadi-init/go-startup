package main

import (
	"gin-app/api"
	"log"
)

func main() {
	// Get the router from the api package
	router := api.SetupRouter()

	// Run the server
	log.Printf("Server starting on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
