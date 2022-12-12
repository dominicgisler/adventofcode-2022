package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day23 struct {
	core.Core
}

func main() {
	d := Day23{}
	d.Init()
	d.part1()
	d.part2()
}
