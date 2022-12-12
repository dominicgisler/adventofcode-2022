package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day16 struct {
	core.Core
}

func main() {
	d := Day16{}
	d.Init()
	d.part1()
	d.part2()
}
