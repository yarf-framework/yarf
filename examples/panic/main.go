package main

import (
	"fmt"
	"github.com/yarf-framework/yarf"
)

// Define a simple resource
type Panic struct {
	yarf.Resource // Extend the yarf.Resource by composition
}

// Implement a GET handler that panics
func (p *Panic) Get(c *yarf.Context) error {
	c.Render("I'm panicking!")

	panic("Totally panicking!")

	return nil
}

func PanicHandler() {
	if err := recover(); err != nil {
		fmt.Printf("Handling panic: %v \n", err)
	}
}

// Entry point of the executable application
// It runs a default server listening on http://localhost:8080
func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Add route/resource
	y.Add("/", new(Panic))

	// Set our custom panic handler
	y.PanicHandler = PanicHandler

	// Start server listening on port 8080
	y.Start(":8080")
}
