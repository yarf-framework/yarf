package main

import (
	"github.com/yarf-framework/yarf"
)

type HelloMiddleware struct {
	yarf.Middleware
}

func (m *HelloMiddleware) PreDispatch(c *yarf.Context) error {
	c.Render("Hello from middleware! \n\n")

	return nil
}

func (m *HelloMiddleware) PostDispatch(c *yarf.Context) error {
	c.Render("\n\nGoodbye from middleware!")

	return nil
}

type ExtraMiddleware struct {
	yarf.Middleware
}

func (m *ExtraMiddleware) PreDispatch(c *yarf.Context) error {
	c.Render("Extra from nested middleware! \n\n")

	return nil
}

func (m *ExtraMiddleware) PostDispatch(c *yarf.Context) error {
	c.Render("\n\nExtra from nested middleware!")

	return nil
}
