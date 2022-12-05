package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day5 struct {
	core.Core
}

func main() {
	d := Day5{}
	d.Init()
	d.part1()
	d.part2()
}
