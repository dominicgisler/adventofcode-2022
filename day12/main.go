package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day12 struct {
	core.Core
	walked      []string
	land        [][]rune
	path        [][]int
	lowestSteps int
}

func main() {
	d := Day12{}
	d.Init()
	d.part1()
	d.part2()

}

func (d *Day12) contains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}
