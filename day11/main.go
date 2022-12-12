package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day11 struct {
	core.Core
}

func main() {
	d := Day11{}
	d.Init()
	d.part1()
	d.part2()
}
