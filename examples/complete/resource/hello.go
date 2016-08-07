package resource

import (
	"github.com/yarf-framework/yarf"
)

// Hello composites yarf.Resource
type Hello struct {
	yarf.Resource
}

// Get implements the GET handler with optional name parameter
func (h *Hello) Get(c *yarf.Context) error {
	name := c.Param("name")

	salute := "Hello"
	if name != "" {
		salute += ", " + name
	}
	salute += "!"

	c.Render(salute)

	return nil
}
