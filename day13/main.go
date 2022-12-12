package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day13 struct {
	core.Core
}

func main() {
	d := Day13{}
	d.Init()
	d.part1()
	d.part2()
}
