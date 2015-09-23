package middleware

import (
	"github.com/yarf-framework/yarf"
)

type Hello struct {
	yarf.Middleware
}

func (m *Hello) PreDispatch() error {
	m.Render("Hello from middleware! \n\n")

	return nil
}

func (m *Hello) PostDispatch() error {
	m.Render("\n\nGoodbye from middleware!")

	return nil
}
