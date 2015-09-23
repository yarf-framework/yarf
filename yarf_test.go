package yarf

import (
	"testing"
)

type MockResource struct {
    Resource
}

type MockMiddleware struct {
    Middleware
}

func TestYarfAdd(t *testing.T) {
    y := New()
    r := new(MockResource)
    
    y.Add("/test", r)
    
    if len(y.routes) != 1 {
        t.Fatalf("Added 1 route, found %d in the list.", len(y.routes))
    }
    if y.routes[0].(*route).path != "/test" {
        t.Fatalf("Added /test path. Found %s", y.routes[0].(*route).path)
    }
    if y.routes[0].(*route).handler != r {
        t.Fatal("Added a handler. Handler found seems to be different")
    }
    
    y.Add("/test/2", r)
    
    if len(y.routes) != 2 {
        t.Fatalf("Added 2 routes, found %d routes in the list.", len(y.routes))
    }
    
    if y.routes[0].(*route).handler != y.routes[1].(*route).handler {
        t.Fatal("Added a handler to 2 routes. Handlers found seems to be different")
    }
}

func TestYarfAddGroup(t *testing.T) {
    y := New()
    g := RouteGroup("/group")
    
    y.AddGroup(g)
    
    if len(y.routes) != 1 {
        t.Fatalf("Added 1 route group, found %d in the list.", len(y.routes))
    }
    if y.routes[0].(*routeGroup).prefix != "/group" {
        t.Fatalf("Added a /group route prefix. Found %s", y.routes[0].(*routeGroup).prefix)
    }
}

func TestYarfInsert(t *testing.T) {
    y := New()
    m := new(MockMiddleware)
    
    y.Insert(m)
    
    if len(y.middleware) != 1 {
        t.Fatalf("Added 1 middleware, found %d in the list.", len(y.routes))
    }
    if y.middleware[0] != m {
        t.Fatal("Added a middleware. Stored one seems to be different")
    }
}
