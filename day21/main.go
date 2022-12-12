package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day21 struct {
	core.Core
}

func main() {
	d := Day21{}
	d.Init()
	d.part1()
	d.part2()
}
