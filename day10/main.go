package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day10 struct {
	core.Core
}

func main() {
	d := Day10{}
	d.Init()
	d.part1()
	d.part2()
}
