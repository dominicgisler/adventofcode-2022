package core

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Core struct {
	Input      string
	InputLines []string
	Debug      bool
}

func (c *Core) Init() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	c.Input = string(data)
	c.InputLines = strings.Split(c.Input, "\n")
}

func (c *Core) StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (c *Core) DebugStr(format string, a ...any) {
	if c.Debug {
		fmt.Printf(format, a)
	}
}
