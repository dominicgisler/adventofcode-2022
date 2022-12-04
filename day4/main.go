package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day4 struct {
	core.Core
}

func main() {
	d := Day4{}
	d.Init()
	d.part1()
	d.part2()
}
