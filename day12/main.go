package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day12 struct {
	core.Core
}

func main() {
	d := Day12{}
	d.Init()
	d.part1()
	d.part2()
}
