package yarf

import (
	"testing"
	"net/url"
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