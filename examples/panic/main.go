package main

import (
	"fmt"
	"github.com/yarf-framework/yarf"
)

// Panic defines a simple resource
type Panic struct {
	yarf.Resource // Extend the yarf.Resource by composition
}

// Get implements a GET handler that panics
func (p *Panic) Get(c *yarf.Context) error {
	c.Render("I'm panicking!")

	panic("Totally panicking!")

	// The next line is unreachable (govet)
	//return nil
}

// PanicHandler is used to catch panics and display the error message
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
