package yarf

import (
	"net/url"
	"testing"
)

// Empty handler
type Handler struct {
	Resource
}

func TestRouterRootMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/", h)

	// Matching routes
	rs := []string{"/", ""}

	// Check
	for _, s := range rs {
		if !r.Match(s, c) {
			t.Errorf("'%s' should match against '/'", s)
		}
	}
}

func TestRouterRootUnmatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/", h)

	// Non-matching routes
	rs := []string{"/something", "something", "/some/thing", "some/thing"}

	// Check
	for _, s := range rs {
		if r.Match(s, c) {
			t.Errorf("'%s' shouldn't match against '/'", s)
		}
	}
}

func TestRouter1LevelMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/level", h)

	// Matching routes
	rs := []string{"/level", "level"}

	// Check
	for _, s := range rs {
		if !r.Match(s, c) {
			t.Errorf("'%s' should match against '/level'", s)
		}
	}
}

func TestRouter1LevelUnmatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/level", h)

	// Non-matching routes
	rs := []string{"/", "", "/:level", "/Level", "/some/thing", "some/thing", "/more/levels/to/be/sure/it/shouldn't/matter/", "/with/trailer/"}

	// Check
	for _, s := range rs {
		if r.Match(s, c) {
			t.Errorf("'%s' shouldn't match against '/level'", s)
		}
	}
}

func TestRouterMultiLevelMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/a/b/c", h)

	// Matching routes
	rs := []string{"/a/b/c", "a/b/c", "/a/b/c/", "a/b/c/"}

	// Check
	for _, s := range rs {
		if !r.Match(s, c) {
			t.Errorf("'%s' should match against '/a/b/c", s)
		}
	}
}

func TestRouterMultiLevelUnmatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/a/b/c", h)

	// Non-matching routes
	rs := []string{"/", "", "/:a/b/c", "/A/B/C", "/some/thing", "some/thing", "/more/levels/to/be/sure/it/shouldn't/matter", "///", "/almost/trailer/"}

	// Check
	for _, s := range rs {
		if r.Match(s, c) {
			t.Errorf("'%s' shouldn't match against '/a/b/c'", s)
		}
	}
}

func TestRouter1LevelParamMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/:param", h)

	// Matching routes
	rs := []string{"/a", "a", "/cafewafewa", "/:paramStyle", "/trailer/"}

	// Check
	for _, s := range rs {
		if !r.Match(s, c) {
			t.Errorf("'%s' should match against '/:param'", s)
		}
	}
}

func TestRouter1LevelParamUnmatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/:param", h)

	// Non-matching routes
	rs := []string{"/", "", "/some/thing", "some/thing", "/more/levels/to/be/sure/it/shouldn't/matter/", "/with/trailer/"}

	// Check
	for _, s := range rs {
		if r.Match(s, c) {
			t.Errorf("'%s' shouldn't match against '/:param'", s)
		}
	}
}

func TestRouterMultiLevelParamMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/a/b/:param", h)

	// Matching routes
	rs := []string{"/a/b/c", "a/b/c", "/a/b/c/", "a/b/c/", "/a/b/:c", "/a/b/:param"}

	// Check
	for _, s := range rs {
		if !r.Match(s, c) {
			t.Errorf("'%s' should match against '/a/b/:param'", s)
		}
	}
}

func TestRouterMultiLevelParamUnmatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create route
	r := Route("/a/b/:param", h)

	// Non-matching routes
	rs := []string{"/", "", "/a/b", "a/b", "/a/b/c/d", "/a/b/"}

	// Check
	for _, s := range rs {
		if r.Match(s, c) {
			t.Errorf("'%s' shouldn't match against '/a/b/:param'", s)
		}
	}
}

func TestRouterGroupMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create group
	g := RouteGroup("/v1")
	g.Add("/test/:param", h)

	// Matching routes
	rs := []string{"/v1/test/test", "/v1/test/:param/"}

	// Check
	for _, s := range rs {
		if !g.Match(s, c) {
			t.Errorf("'%s' should match", s)
		}
	}
}

func TestRouterGroupNotMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create group
	g := RouteGroup("/v1")
	g.Add("/test/:param", h)

	// Non-Matching routes
	rs := []string{"/test/test", "/v1/test", "/v1/test/a/b", "/v1", "/"}

	// Check
	for _, s := range rs {
		if g.Match(s, c) {
			t.Errorf("'%s' shouldn't match", s)
		}
	}
}

func TestRouterNestedGroupMatch(t *testing.T) {
	// Create empty handler
	h := new(Handler)

	// Create empty context
	c := new(Context)
	c.Params = url.Values{}

	// Create groups
	l1 := RouteGroup("/level1")
	l2 := RouteGroup("/level2")
	l3 := RouteGroup("/level3")

	// Add one route
	l3.Add("/test/:param", h)

	// Neste into:
	// - /level1/level2/level3/test/:param
	// - /level2/level3/test/:param
	// - /level3/test/:param
	l2.AddGroup(l3)
	l1.AddGroup(l2)

	// Level 3 matching routes
	rs := []string{"/level3/test/test", "/level3/test/:param/"}

	// Check
	for _, s := range rs {
		if !l3.Match(s, c) {
			t.Errorf("'%s' should match", s)
		}
	}

	// Level 2 matching routes
	rs = []string{"/level2/level3/test/test", "/level2/level3/test/:param/"}

	// Check
	for _, s := range rs {
		if !l2.Match(s, c) {
			t.Errorf("'%s' should match", s)
		}
	}

	// Level 1 matching routes
	rs = []string{"/level1/level2/level3/test/test", "/level1/level2/level3/test/:param/"}

	// Check
	for _, s := range rs {
		if !l1.Match(s, c) {
			t.Errorf("'%s' should match", s)
		}
	}
}

func BenchmarkRouteMatch_short(b *testing.B) {
	h := &Handler{}
	c := &Context{}
	r := Route("/test", h)
	for i := 0; i < b.N; i++ {
		r.Match("/test", c)
		r.Match("/nomatch", c)
	}
}

func BenchmarkRouteMatch_long(b *testing.B) {
	h := &Handler{}
	c := &Context{}
	routeString := "/very/long/route/with/ten/separate/parts/eight/nine/ten"
	r := Route(routeString, h)
	for i := 0; i < b.N; i++ {
		r.Match(routeString, c)
		r.Match("/short/request/url", c)
		r.Match("/very/long/route/with/ten/separate/parts/that/doesnt/match", c)
	}
}

func BenchmarkRouteMatch_emptyParts(b *testing.B) {
	h := &Handler{}
	c := &Context{}
	r := Route("/route///with//lots////of///empty///parts/", h)
	for i := 0; i < b.N; i++ {
		r.Match("/route///with/lots/of////empty////parts/", c)
		r.Match("/request/////url/////////////with//////////////tons//of/empty///////////parts/////////////test", c)
	}
}

func BenchmarkRouteGroupMatch_short(b *testing.B) {
	h := &Handler{}
	c := &Context{}
	r := RouteGroup("/prefix")
	r.Add("/suffix", h)
	for i := 0; i < b.N; i++ {
		r.Match("/test", c)
		r.Match("/nomatch", c)
	}
}

func BenchmarkRouteGroupMatch_nested(b *testing.B) {
	h := &Handler{}
	c := &Context{}
	// create a set of nested RouteGroups 20 levels deep
	g := RouteGroup("/test")
	g.Add("/router", h)
	path := "/test/router"
	for i := 0; i < 19; i++ {
		r := RouteGroup("/test")
		r.AddGroup(g)
		g = r
		path = "/test" + path
	}
	for i := 0; i < b.N; i++ {
		g.Match(path, c)
		g.Match(path+"matchfail", c)
	}
}
