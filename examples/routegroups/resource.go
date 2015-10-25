package main

import (
	"github.com/yarf-framework/yarf"
)

type Hello struct {
	yarf.Resource
}

// Implement the GET handler with optional name parameter
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

type HelloV2 struct {
	yarf.Resource
}

func (h *HelloV2) Get(c *yarf.Context) error {
	name := c.Param("name")

	salute := "(v2) Hello"
	if name != "" {
		salute += ", " + name
	}
	salute += "!"

	c.Render(salute)

	return nil
}
