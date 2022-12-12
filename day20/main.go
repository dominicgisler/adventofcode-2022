package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day20 struct {
	core.Core
}

func main() {
	d := Day20{}
	d.Init()
	d.part1()
	d.part2()
}
