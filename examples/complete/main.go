// The complete example implements all possible features from YARF.
package main

import (
	"github.com/yarf-framework/extras/logger"
	"github.com/yarf-framework/yarf"
	"github.com/yarf-framework/yarf/examples/complete/middleware"
	"github.com/yarf-framework/yarf/examples/complete/resource"
)

func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Create resource
	hello := new(resource.Hello)

	// Add resource to multiple routes
	y.Add("/", hello)
	y.Add("/hello/:name", hello)

	// Create /extra route group
	e := yarf.RouteGroup("/extra")

	// Add custom middleware to /extra
	e.Insert(new(middleware.Hello))

	// Add same routes to /extra group
	e.Add("/", hello)
	e.Add("/hello/:name", hello)

	// Save group
	y.AddGroup(e)

	// Add logger middleware at the end of the chain
	y.Insert(new(logger.Logger))

	// Start server listening on port 8080
	y.Start(":8080")
}
