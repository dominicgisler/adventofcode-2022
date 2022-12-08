package main

import (
	"fmt"
)

func (d *Day8) part2() {

	var trees [][]rune
	for y, line := range d.InputLines {
		trees = append(trees, []rune{})
		for _, r := range line {
			trees[y] = append(trees[y], r)
		}
	}

	var bestScore int
	for y, l := range trees {
		for x, t := range l {
			score := 1
			var cnt int

			cnt = 0
			for i := y - 1; i >= 0; i-- {
				if trees[i][x] < t {
					cnt++
				}
				if trees[i][x] >= t {
					cnt++
					i = -1
				}
			}
			score *= cnt

			cnt = 0
			for i := y + 1; i < len(trees); i++ {
				if trees[i][x] < t {
					cnt++
				}
				if trees[i][x] >= t {
					cnt++
					i = len(trees)
				}
			}
			score *= cnt

			cnt = 0
			for i := x - 1; i >= 0; i-- {
				if trees[y][i] < t {
					cnt++
				}
				if trees[y][i] >= t {
					cnt++
					i = -1
				}
			}
			score *= cnt

			cnt = 0
			for i := x + 1; i < len(l); i++ {
				if trees[y][i] < t {
					cnt++
				}
				if trees[y][i] >= t {
					cnt++
					i = len(l)
				}
			}
			score *= cnt

			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Println("best score:", bestScore)
}
