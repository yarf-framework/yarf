[![GoDoc](https://godoc.org/github.com/yarf-framework/yarf?status.svg)](https://godoc.org/github.com/yarf-framework/yarf)
[![Build Status](https://travis-ci.org/yarf-framework/yarf.svg?branch=master)](https://travis-ci.org/yarf-framework/yarf)

# YARF: Yet Another REST Framework

YARF is a micro-framework designed to build REST APIs and web services in a fast and simple way. 
It can be used to develop any kind of web application extending some features, but that's not the purpose of the framework.
Designed after Go's composition features, takes a new approach to write simple and DRY code.


## Version

Current version is: **0.2**


## Documentation

https://godoc.org/github.com/yarf-framework/yarf


## Code

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
    y := yarf.Yarf()
    y.Add("/", new(Hello))
    y.Start(":8080")
}

```

For more code and examples demonstrating all YARF features, please refer to the 'examples' directory.


## Features

#### v0.2

- **Resource based design:** Each resource can implement one, several or all HTTP methods needed (GET, POST, DELETE, etc.).
- **Simple router:** Matches exact URLs against resources for increased performance. The routes supports params in the form /:param.
- **Optional parameters:** Supported using multiple routes on the same Resource.
- **Middleware:** Support at router level, all routes will be pre-filtered and post-filtered by Middleware handlers.
- **Route groups:** Routes can be grouped into a route prefix and handle their own middleware.
- **Nested groups:** As routes can be grouped into a route prefix, other groups can be also grouped allowing for nested prefixes and middleware layers.



## Future

Features that are planned to be implemented soon.

- Framework Tests.
- Framework Benchmarks.
- Increase custom errors collection.
- Middleware support at Resource level.
- Middleware and Lib packages with extra/useful functionality to add to any web application.
- Gzip responses.


## Contribute

**Yes, please!**

- Use and test YARF and/or packages included.
- Implement new web applications based on Yarf.
- Report issues/bugs/comments/suggestions on Github
- Fork and send me your pull requests with descriptions of modifications/new features

