package main

import (
	"github.com/vincentJunior1/test-kriya/models"
	"github.com/vincentJunior1/test-kriya/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	r := routes.SetupRouter()

	// Initialize database
	models.SetupDatabase()

	// Start the HTTP API
	r.Run()
}
