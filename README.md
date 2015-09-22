[![GoDoc](https://godoc.org/github.com/yarf-framework/yarf?status.svg)](https://godoc.org/github.com/yarf-framework/yarf)
[![Build Status](https://travis-ci.org/yarf-framework/yarf.svg?branch=master)](https://travis-ci.org/yarf-framework/yarf)


# YARF: Yet Another REST Framework

YARF is a micro-framework designed to build REST APIs and web services in a fast and simple way. 
It can be used to develop any kind of web application extending some features, but that's not the purpose of the framework.
Designed after Go's composition features, takes a new approach to write simple and DRY code.



## Documentation

https://godoc.org/github.com/yarf-framework/yarf



## Getting started

Here's a transcription from our examples/simple package. 
This is a very simple Hello World web application example. 


```go
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
    y := yarf.New()
    
    y.Add("/", new(Hello))
    
    y.Start(":8080")
}

```

For more code and examples demonstrating all YARF features, please refer to the 'examples' directory.



## Features


### Struct composition based design

YARF resources are custom structs that act as handlers for HTTP methods. 
Each resource can implement one, several or all HTTP methods needed (GET, POST, DELETE, etc.).
Resources are created using Go's struct composition and you only have to declare the yarf.Resource type into your own struct.

Example: 

```go

import "github.com/yarf-framework/yarf"

// Define a simple resource
type Hello struct {
    yarf.Resource
}

// Implement the GET method
func (h *Hello) Get() error {
    h.Render("Hello world!")
    
    return nil
}

```


### Simple router

Using a strict match model, it matches exact URLs against resources for increased performance and clarity during routing. 
The routes supports parameters in the form '/:param'.

The route: 

```go
/hello/:name
```

Will match: 

```
/hello/Joe
/hello/nobody
/hello/somebody/
/hello/:name
/hello/etc
```

But it won't match: 

```
/
/hello
/hello/Joe/AndMark
/Joe/:name
/any/thing
```

You can define optional parameters using multiple routes on the same Resource.


### Middleware support

Middleware support is implemented in a similar way as Resources, by using composition.  
Routes will be pre-filtered and post-filtered by Middleware handlers when they're inserted in the router.

Example: 

```go
import "github.com/yarf-framework/yarf"

// Define your middleware and composite yarf.Middleware
type HelloMiddleware struct {
    yarf.Middleware
}

// Implement only the PreDispatch method, PostDispatch not needed.
func (m *HelloMiddleware) PreDispatch() error {
    m.Render("Hello from middleware! \n") // Render to response.

    return nil
}

// Insert your middlewares to the server
func main() {
    y := yarf.New()

    // Add middleware
    y.Insert(new(HelloMiddleware))
    
    // Define routes
    // ...
    // ...
    
    // Start the server
    y.Start()
}
```


### Route groups

Routes can be grouped into a route prefix and handle their own middleware.


### Nested groups

As routes can be grouped into a route prefix, other groups can be also grouped allowing for nested prefixes and middleware layers.

Example: 

```go
import "github.com/yarf-framework/yarf"

// Entry point of the executable application
// It runs a default server listening on http://localhost:8080
//
// URLs after configuration:
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

    // Add main resource to multiple routes at root level.
    y.Add("/", hello)
    y.Add("/hello/:name", hello)

    // Create /v2 prefix route group
    g := yarf.RouteGroup("/v2")

    // Add /v2/ routes to the group
    g.Add("/", hellov2)
    g.Add("/hello/:name", hellov2)

    // Use middleware only on the /v2/ group
    g.Insert(new(HelloMiddleware))

    // Add group to Yarf routes
    y.AddGroup(g)

    // Create another group for nesting into it.
    n := yarf.RouteGroup("/extra")

    // Nest /v2 group into /extra/v2
    n.AddGroup(g)

    // Use another middleware only for this /extra/v2 group
    n.Insert(new(ExtraMiddleware))

    // Add group to Yarf
    y.AddGroup(n)

    // Start server listening on port 8080
    y.Start(":8080")
}
```

Check the ./examples/routegroups demo for the complete working implementation.


### Route caching

A route cache is enabled by default to improve dispatch speed, but sacrificing memory space. 
If you're running out of RAM memory and/or your app has too many possible routes that may not fit, you should disable the route cache.

To enable/disable the route cache, just set the UseCache flag of the Yarf object: 

```go
y := yarf.New()
y.UseCache = false
```


## Performance

On initial benchmarks, the framework seems to perform very well compared with other similar frameworks. 
Even when there are faster frameworks, under high load conditions and thanks to the route caching method, 
YARF seems to perform as good or even better than the fastests that work better under simpler conditions.

Check the benchmarks repository to run your own:

[https://github.com/yarf-framework/benchmarks](https://github.com/yarf-framework/benchmarks)



## HTTPS support

Support for running HTTPS server from the net/http package. 

### Using the default server

```go
func main() {
    y := yarf.New()
    
    // Setup the app
    // ...
    // ...
    
    // Start https listening on port 443
    y.StartTLS(":443", certFile, keyFile)
}

```

### Using a custom server

```go
func main() {
    y := yarf.New()

    // Setup the app
    // ...
    // ...

    // Configure custom http server and set the yarf object as Handler.
    s := &http.Server{
        Addr:           ":443",
        Handler:        y,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServeTLS(certFile, keyFile)
}

```