/*
Package yarf (Yet Another Rest Framework) provides the foundations to write REST APIs in a fast and simple way.

This is a full YARF "Hello World" example: 

	package main
	
	import (
	    "github.com/yarf-framework/yarf"
	)
	
	// Define a simple resource
	type Hello struct {
	    yarf.Resource
	}
	
	// Implement the GET method
	func (h *Hello) Get() error {
	    h.Render("Hello world!")
	    
	    return nil
	}
	
	// Run app server on http://localhost:8080
	func main() {
	    y := yarf.Yarf()
	    y.Add("/", new(Hello))
	    y.Start(":8080")
	}
	


Basic features: 

- Resource based design: Each resource can implement one, several or all HTTP methods needed (GET, POST, DELETE, etc.).


- Simple router: Matches exact URLs against resources for increased performance. The routes supports params in the form /:param.


- Optional parameters: Supported using multiple routes on the same Resource.


- Middleware: Support at router level, all routes will be pre-filtered and post-filtered by Middleware handlers.


- Route groups: Routes can be grouped into a route prefix and handle their own middleware.


- Nested groups: As routes can be grouped into a route prefix, other groups can be also grouped allowing for nested prefixes and middleware layers.


*/