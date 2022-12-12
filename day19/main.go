package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day19 struct {
	core.Core
}

func main() {
	d := Day19{}
	d.Init()
	d.part1()
	d.part2()
}
