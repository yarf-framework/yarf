package main

import (
	"github.com/yarf-framework/yarf"
)

// Hello defines a simple resource by compositing yarf.Resource
type Hello struct {
	yarf.Resource
}

// Get implements the GET handler
func (h *Hello) Get(c *yarf.Context) error {
	c.Render("Hello world!")

	return nil
}

// Entry point of the executable application
// It runs a default server listening on http://localhost:8080
func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Add route/resource
	y.Add("/", new(Hello))

	// Start server listening on port 8080
	y.Start(":8080")
}
