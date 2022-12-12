package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day14 struct {
	core.Core
}

func main() {
	d := Day14{}
	d.Init()
	d.part1()
	d.part2()
}
