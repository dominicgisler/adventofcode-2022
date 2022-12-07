package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day7 struct {
	core.Core
}

func main() {
	d := Day7{}
	d.Init()
	d.part1()
	d.part2()
}
