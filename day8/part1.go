package main

import (
	"fmt"
)

func (d *Day8) part1() {

	var trees [][]rune
	for y, line := range d.InputLines {
		trees = append(trees, []rune{})
		for _, r := range line {
			trees[y] = append(trees[y], r)
		}
	}

	var cnt int
	for y, l := range trees {
		for x, t := range l {
			if y == 0 || y == len(trees)-1 || x == 0 || x == len(l)-1 {
				cnt++
				continue
			}

			visible := true

			for i := 0; i < y && visible; i++ {
				if trees[i][x] >= t {
					visible = false
				}
			}
			if visible {
				cnt++
				continue
			}

			visible = true
			for i := y + 1; i < len(trees) && visible; i++ {
				if trees[i][x] >= t {
					visible = false
				}
			}
			if visible {
				cnt++
				continue
			}

			visible = true
			for i := 0; i < x && visible; i++ {
				if trees[y][i] >= t {
					visible = false
				}
			}
			if visible {
				cnt++
				continue
			}

			visible = true
			for i := x + 1; i < len(l) && visible; i++ {
				if trees[y][i] >= t {
					visible = false
				}
			}
			if visible {
				cnt++
				continue
			}
		}
	}

	fmt.Println("visible:", cnt)
}
