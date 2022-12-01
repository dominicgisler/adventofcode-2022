package core

import (
	"os"
	"strings"
)

type Core struct {
	Input      string
	InputLines []string
}

func (c *Core) Init() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	c.Input = string(data)
	c.InputLines = strings.Split(c.Input, "\n")
}
