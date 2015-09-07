package yarf

import (
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
