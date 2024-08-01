package main

import (
	"tskmgr/router"
)

func main() {
	// Initialize the router and set up routes
	r := router.Route()

	// Start the server on port 8080
	r.Run("localhost:8080")
}
