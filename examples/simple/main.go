package simple

import (
	"github.com/leonelquinteros/yarf"
)

// Define a simple resource
type Hello struct {
	yarf.Resource // Extend the yarf.Resource by composition
}

// Implement the GET handler
// All YARF handlers receive a single *yarf.Context param which wraps the entire request data.
func (h *Hello) Get() error {
	// Render a string to the response
	h.Render("Hello world!")

	// Terminate execution without errors
	return nil
}

// Entry point of the executable application
func main() {
	// Create and start a new empty YARF server
	server, err := yarf.Server()
	if err != nil {
		// For the purposes of the example, we'll panic on server creation error.
		// You may want to implement your own/convenient error handling.
		panic(err.Error())
	}

	// Add route/resource
	server.Add("/", new(Hello))
}
