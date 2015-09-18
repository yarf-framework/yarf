// The complete example implements all possible features from YARF.
package main

import (
	"github.com/yarf-framework/yarf"
	"github.com/yarf-framework/yarf/examples/complete/resource"
	"github.com/yarf-framework/yarf/lib/middleware"
)

func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Create resource
	hello := new(resource.Hello)

	// Add resource to multiple routes
	y.Add("/", hello)
	y.Add("/hello/:name", hello)

	// Add gzip middleware first in the chain.
	y.Insert(new(middleware.Gzip))

	// Add logger middleware at the end of the chain
	y.Insert(new(middleware.Logger))

	// Start server listening on port 8080
	y.Start(":8080")
}
