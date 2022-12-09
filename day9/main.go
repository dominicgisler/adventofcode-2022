package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day9 struct {
	core.Core
}

func main() {
	d := Day9{}
	d.Init()
	d.part1()
	d.part2()
}
