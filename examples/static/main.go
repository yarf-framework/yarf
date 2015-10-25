// static example package demonstrates how to easily implement
// a yarf handler that serves static files using the net/http package.
package main

import (
	"github.com/yarf-framework/yarf"
	"net/http"
)

// Define a simple resource
type Static struct {
	yarf.Resource // Extend the yarf.Resource by composition

	path string // Directory to serve static files from.
}

// Implement the static files handler
func (s *Static) Get(c *yarf.Context) error {
	http.FileServer(http.Dir(s.path)).ServeHTTP(c.Response, c.Request)

	return nil
}

// StaticDir constructs a Static handler and sets the path to serve under the route.
func StaticDir(path string) *Static {
	s := new(Static)
	s.path = path

	return s
}

// Entry point of the executable application
// It runs a default server listening on http://localhost:8080
func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Add routes/resources
	y.Add("/", StaticDir("/tmp"))
	y.Add("/test", StaticDir("/var/www/test"))

	// Start server listening on port 8080
	y.Start(":8080")
}
