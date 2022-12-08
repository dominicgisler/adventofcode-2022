package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day8 struct {
	core.Core
}

func main() {
	d := Day8{}
	d.Init()
	d.part1()
	d.part2()
}
