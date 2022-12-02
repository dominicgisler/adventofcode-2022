package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day2 struct {
	core.Core
}

func main() {
	d := Day2{}
	d.Init()
	d.part1()
	d.part2()
}
