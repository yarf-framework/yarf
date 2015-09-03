# YARF: Yet Another REST Framework

YARF is a micro-framework designed to build REST APIs and web services in a fast and simple way. 
It can be used to develop any kind of web application extending some features, but that's not the purpose of the framework.
Designed after Go's composition features, takes a new approach to write simple and DRY code.


## Version

Current version is: *0.1b*

The project is still in development and probably not working, but growing step by step into a functional framework.
The first release of the version 0.1 should be a working framework with a tiny set of basic features.


## Code

Here's a transcription from our examples/simple package. 
This is a very simple Hello World web application example. 


```go
package main

import (
    "github.com/leonelquinteros/yarf"
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

#### v0.1

- Resource based design, each resource can implement one or all HTTP methods needed (GET, POST, DELETE, etc.).
- Simple router matches exact URLs against resources. The routes supports params in the form /:param.
- Support for optional parameters defining multiple routes against the same Resource.
- Middleware support at router level. All routes (and methods) will be pre-filtered and post-filtered by Middleware handlers.



## Future

Features that are planned to be implemented soon.

- Framework Tests
- Framework Benchmarks
- Middleware support at Resource level
- Middleware and Lib packages with extra/useful functionality to add to any web application.


## Contribute

*Yes, please!*

- Use and test YARF and/or packages included.
- Implement new web applications based on Yarf.
- Report issues/bugs/comments/suggestions on Github
- Fork and send me your pull requests with descriptions of modifications/new features

