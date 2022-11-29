package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day1 struct {
	core.Core
}

func main() {
	d := Day1{}
	d.Init()
	d.part1()
	d.part2()
}
