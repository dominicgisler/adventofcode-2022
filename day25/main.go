package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day25 struct {
	core.Core
}

func main() {
	d := Day25{}
	d.Init()
	d.part1()
	d.part2()
}
