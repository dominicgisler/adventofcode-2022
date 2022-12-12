package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day24 struct {
	core.Core
}

func main() {
	d := Day24{}
	d.Init()
	d.part1()
	d.part2()
}
