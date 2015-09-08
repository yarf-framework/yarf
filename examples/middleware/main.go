package main

import (
	"github.com/yarf-framework/yarf"
)

// Entry point of the executable application
// It runs a default server listening on http://localhost:8080
func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Add middleware
	y.Insert(new(HelloMiddleware))

	// Create resource
	hello := new(Hello)

	// Add resource to multiple routes
	y.Add("/", hello)
	y.Add("/hello/:name", hello)

	// Start server listening on port 8080
	y.Start(":8080")
}
