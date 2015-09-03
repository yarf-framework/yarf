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

// PostDispatch includes code to be executed after every Resource request.
func (m *HelloMiddleware) PostDispatch() error {
	m.Render("\n\nGoodbye from middleware!")

	return nil
}
