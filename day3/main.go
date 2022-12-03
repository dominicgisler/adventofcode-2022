package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day3 struct {
	core.Core
}

func main() {
	d := Day3{}
	d.Init()
	d.part1()
	d.part2()
}

func (d *Day3) contains(s []rune, str rune) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (d *Day3) getPrio(r rune) int {

	const lowerDiff = -96
	const upperDiff = -38

	if d.contains([]rune("abcdefghijklmnopqrstuvwxyz"), r) {
		return int(r + lowerDiff)
	} else {
		return int(r + upperDiff)
	}
}
