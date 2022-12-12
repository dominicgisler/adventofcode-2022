package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day17 struct {
	core.Core
}

func main() {
	d := Day17{}
	d.Init()
	d.part1()
	d.part2()
}
