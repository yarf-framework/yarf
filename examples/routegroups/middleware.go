package main

import (
	"github.com/yarf-framework/yarf"
)

type HelloMiddleware struct {
	yarf.Middleware
}

func (m *HelloMiddleware) PreDispatch() error {
	m.Render("Hello from middleware! \n\n")

	return nil
}

func (m *HelloMiddleware) PostDispatch() error {
	m.Render("\n\nGoodbye from middleware!")

	return nil
}

type ExtraMiddleware struct {
	yarf.Middleware
}

func (m *ExtraMiddleware) PreDispatch() error {
	m.Render("Extra from nested middleware! \n\n")

	return nil
}

func (m *ExtraMiddleware) PostDispatch() error {
	m.Render("\n\nExtra from nested middleware!")

	return nil
}
