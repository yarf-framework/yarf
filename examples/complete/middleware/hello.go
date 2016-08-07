package middleware

import (
	"github.com/yarf-framework/yarf"
)

// Hello composites yarf.Middleware
type Hello struct {
	yarf.Middleware
}

// PreDispatch renders a hardcoded string
func (m *Hello) PreDispatch(c *yarf.Context) error {
	c.Render("Hello from middleware! \n\n")

	return nil
}

// PostDispatch renders a hardcoded string
func (m *Hello) PostDispatch(c *yarf.Context) error {
	c.Render("\n\nGoodbye from middleware!")

	return nil
}
