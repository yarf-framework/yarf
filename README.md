[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GitHub version](https://badge.fury.io/gh/yarf-framework%2Fyarf.svg)](https://badge.fury.io/gh/yarf-framework%2Fyarf)
[![GoDoc](https://godoc.org/github.com/yarf-framework/yarf?status.svg)](https://godoc.org/github.com/yarf-framework/yarf)
[![Build Status](https://travis-ci.org/yarf-framework/yarf.svg?branch=master)](https://travis-ci.org/yarf-framework/yarf)
[![codecov](https://codecov.io/gh/yarf-framework/yarf/branch/master/graph/badge.svg)](https://codecov.io/gh/yarf-framework/yarf)
[![Go Report Card](https://goreportcard.com/badge/github.com/yarf-framework/yarf)](https://goreportcard.com/report/github.com/yarf-framework/yarf)
[![Join the chat at https://gitter.im/jinzhu/gorm](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/yarf-framework/yarf) 


# YARF: Yet Another REST Framework

YARF is a fast micro-framework designed to build REST APIs and web services in a fast and simple way. 
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
func (h *Hello) Get(c *yarf.Context) error {
    c.Render("Hello world!")
    
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
func (h *Hello) Get(c *yarf.Context) error {
    c.Render("Hello world!")
    
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


### Route parameters

At this point you know how to define parameters in your routes using the /:param naming convention. 
Now you'll see how easy is to get these parameters by their name from your resources using the Context.Param() method. 

Example: 

For the route: 

```
/hello/:name
```

You can have this resource:

```go
import "github.com/yarf-framework/yarf"

type Hello struct {
    yarf.Resource
}

func (h *Hello) Get(c *yarf.Context) error {
    name := c.Param("name")

    c.Render("Hello, " + name)

    return nil
}

```


### Route wildcards

When some extra freedom is needed on your routes, you can use a `*` as part of your routes to match anything where the wildcard is present. 

The route: 

```
/something/*/here
```

Will match the routes

```
/something/is/here
/something/happen/here
/something/isnt/here
/something/was/here
```

And so on... 

You can also combine this with parameters inside the routes for extra complexity.


### Catch-All wildcard

When using the `*` at the end of any route, the router will match everything from the wildcard and forward. 

The route: 

```
/match/from/here/*
```

Will match: 

```
/match/from/here
/match/from/here/please
/match/from/here/and/forward
/match/from/here/and/forward/for/ever/and/ever
```

And so on...


#### Note about the wildcard

The `*` can only be used by itself and it doesn't works for single character matching like in regex. 

So the route:
 
```
/match/some*
```

Will **NOT** match:

```
/match/some
/match/something
/match/someone
/match/some/please
```


### Context

The Context object is passed as a parameter to all Resource methods and contains all the information related to the ongoing request. 

Check the Context docs for a reference of the object: [https://godoc.org/github.com/yarf-framework/yarf#Context](https://godoc.org/github.com/yarf-framework/yarf#Context)



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
func (m *HelloMiddleware) PreDispatch(c *yarf.Context) error {
    c.Render("Hello from middleware! \n") // Render to response.

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


### Chain and extend

Just use the Yarf object as any http.Handler on a chain. 
Set another http.Handler on the Yarf.Follow property to be followed in case this Yarf router can't match the request. 

Here's an example on how to follow the request to a public file server: 

```go
package main 

import (
    "github.com/yarf-framework/yarf"
    "net/http"
)

func main() {
    y := yarf.New()

    // Add some routes
    y.Add("/hello/:name", new(Hello))
    
    //... more routes here
    
    // Follow to file server
    y.Follow = http.FileServer(http.Dir("/var/www/public"))
    
    // Start the server
    y.Start(":8080")
}
```


### Custom NotFound error handler

You can handle all 404 errors returned by any resource/middleware during the request flow of a Yarf server. 
To do so, you only have to implement a function with the func(c *yarf.Context) signature and set it to your server's Yarf.NotFound property.

```go
y := yarf.New()

// ...

y.NotFound = func(c *yarf.Context) {
    c.Render("This is a custom Not Found handler")
}

// ...
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


## Why another micro-framework? 

Why not?

No, seriously, i've researched for small/fast frameworks in the past for a Go project that I was starting. 
I found several options, but at the same time, none seemed to suit me. 
Some of them make you write weird function wrappers to fit the net/http package style. 
Actually, most of them seem to be function-based handlers. 
While that's not wrong, I feel more comfortable with the resource-based design, and this I also feel aligns better with the spirit of REST. 

In Yarf you create a resource struct that represents a REST resource and it has all HTTP methods available. 
No need to create different routes for GET/POST/DELETE/etc. methods. 

By using composition, you don't need to wrap functions inside functions over and over again to implement simple things like middleware or extension to your methods. 
You can abuse composition to create a huge OO-like design for your business model without sacrifying performance and code readability. 
 
Even while the code style differs from the net/http package, the framework is fully compatible with it and couldn't run without it. 
Extensions and utilities from other frameworks or even the net/http package can be easily implemented into Yarf by just wraping them up into a Resource, 
just as you would do on any other framework by wrapping functions.
  
Context handling also shows some weird designs across frameworks. Some of them rely on reflection to receive any kind of handlers and context types. 
Others make you receive some extra parameter in the handler function that actually brokes the net/http compatibility, and you have to carry that context parameter through all middleware/handler-wrapper functions just to make it available. 
In Yarf, the Context is automatically sent as a parameter to all Resource methods by the framework. 

For all the reasons above, among some others, there is a new framework in town. 
