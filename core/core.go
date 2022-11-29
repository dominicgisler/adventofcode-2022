package core

import (
	"os"
)

type Core struct {
	input string
}

func (c *Core) Init() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	c.input = string(data)
}
