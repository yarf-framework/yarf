package middleware

import (
	"github.com/yarf-framework/yarf"
)

type Hello struct {
	yarf.Middleware
}

func (m *Hello) PreDispatch(c *yarf.Context) error {
	c.Render("Hello from middleware! \n\n")

	return nil
}

func (m *Hello) PostDispatch(c *yarf.Context) error {
	c.Render("\n\nGoodbye from middleware!")

	return nil
}
