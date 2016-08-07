package main

import (
	"github.com/yarf-framework/yarf"
)

// HelloMiddleware composites yarf.Middleware
type HelloMiddleware struct {
	yarf.Middleware
}

// PreDispatch renders a hardcoded string
func (m *HelloMiddleware) PreDispatch(c *yarf.Context) error {
	c.Render("Hello from middleware! \n\n")

	return nil
}

// PostDispatch renders a hardcoded string
func (m *HelloMiddleware) PostDispatch(c *yarf.Context) error {
	c.Render("\n\nGoodbye from middleware!")

	return nil
}

// ExtraMiddleware also composites yarf.Middleware
type ExtraMiddleware struct {
	yarf.Middleware
}

// PreDispatch renders a hardcoded string
func (m *ExtraMiddleware) PreDispatch(c *yarf.Context) error {
	c.Render("Extra from nested middleware! \n\n")

	return nil
}

// PostDispatch renders a hardcoded string
func (m *ExtraMiddleware) PostDispatch(c *yarf.Context) error {
	c.Render("\n\nExtra from nested middleware!")

	return nil
}
