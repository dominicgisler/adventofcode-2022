package main

import (
	"fmt"
)

func (d *Day12) part1() {

	d.land = [][]rune{}
	for y, line := range d.InputLines {
		d.land = append(d.land, []rune{})
		for x, r := range line {
			d.land[y] = append(d.land[y], r)
			if r == 'S' {
				d.path = [][]int{{y, x, 0}}
			}
		}
	}

	var walked []string
	for len(d.path) > 0 {
		y := d.path[0][0]
		x := d.path[0][1]
		ds := d.path[0][2]
		d.path = append(d.path[:0], d.path[1:]...)
		if d.contains(walked, fmt.Sprintf("%d_%d", y, x)) {
			continue
		}
		walked = append(walked, fmt.Sprintf("%d_%d", y, x))
		curr := d.land[y][x]
		if curr == 'S' {
			curr = 'a'
		}
		if curr == 'E' {
			fmt.Println("path", ds) // 447
			d.path = [][]int{}
		}
		for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			ny := y + dir[0]
			nx := x + dir[1]
			if 0 <= ny && ny < len(d.land) && 0 <= nx && nx < len(d.land[0]) {
				next := d.land[ny][nx]
				if next == 'E' {
					next = 'z'
				}
				if next <= curr+1 {
					d.path = append(d.path, []int{ny, nx, ds + 1})
				}
			}
		}
	}
}
