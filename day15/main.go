package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day15 struct {
	core.Core
}

func main() {
	d := Day15{}
	d.Init()
	d.part1()
	d.part2()
}
