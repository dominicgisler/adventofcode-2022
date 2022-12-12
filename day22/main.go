package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day22 struct {
	core.Core
}

func main() {
	d := Day22{}
	d.Init()
	d.part1()
	d.part2()
}
