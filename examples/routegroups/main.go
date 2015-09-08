package main

import (
	"github.com/yarf-framework/yarf"
)

// Entry point of the executable application
// It runs a default server listening on http://localhost:8080
//
// URLs available:
// http://localhost:8080
// http://localhost:8080/hello/:name
// http://localhost:8080/v2
// http://localhost:8080/v2/hello/:name
// http://localhost:8080/extra/v2
// http://localhost:8080/extra/v2/hello/:name
//
func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Create resources
	hello := new(Hello)
	hellov2 := new(HelloV2)

	// Add main resource to multiple routes
	y.Add("/", hello)
	y.Add("/hello/:name", hello)

	// Create /v2 route group
	g := yarf.RouteGroup("/v2")

	// Add v2 routes to the group
	g.Add("/", hellov2)
	g.Add("/hello/:name", hellov2)

	// Use middleware only on the group
	g.Insert(new(HelloMiddleware))

	// Add group to Yarf routes
	y.AddGroup(g)

	// Create another group for nesting
	n := yarf.RouteGroup("/extra")

	// Nest /v2 group into /extra
	n.AddGroup(g)

	// Use another middleware for this group
	n.Insert(new(ExtraMiddleware))

	// Add group to Yarf
	y.AddGroup(n)

	// Start server listening on port 8080
	y.Start(":8080")
}
