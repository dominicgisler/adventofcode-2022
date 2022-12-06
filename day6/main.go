package main

import (
	"github.com/dominicgisler/adventofcode-2022/core"
	"strings"
)

type Day6 struct {
	core.Core
}

func main() {
	d := Day6{}
	d.Init()
	d.part1()
	d.part2()
}

func (d *Day6) uniqueBuffer(buffer []rune) bool {
	var tmp string
	for _, l := range buffer {
		if strings.Contains(tmp, string(l)) {
			return false
		}
		tmp += string(l)
	}
	return true
}

func (d *Day6) getPosition(uniqueNum int) int {
	var pos int
	var buffer []rune
	for i, l := range d.Input {
		if pos > 0 {
			continue
		}
		buffer = append(buffer, l)
		if len(buffer) > uniqueNum {
			buffer = append(buffer[:0], buffer[0+1:]...)
		}
		if len(buffer) == uniqueNum && d.uniqueBuffer(buffer) {
			pos = i
		}
	}
	pos++
	return pos
}
