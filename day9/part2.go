package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (d *Day9) part2() {

	pos := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

	visited := map[string]int{}
	for _, line := range d.InputLines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		num, _ := strconv.Atoi(parts[1])

		for i := 0; i < num; i++ {
			switch dir {
			case "R":
				pos[0][0]++
			case "L":
				pos[0][0]--
			case "U":
				pos[0][1]++
			case "D":
				pos[0][1]--
			}

			for i := 0; i < len(pos)-1; i++ {
				if pos[i][0] > pos[i+1][0]+1 || pos[i][0] < pos[i+1][0]-1 {
					if pos[i][0] > pos[i+1][0]+1 {
						pos[i+1][0]++
					} else if pos[i][0] < pos[i+1][0]-1 {
						pos[i+1][0]--
					}
					if pos[i][1] > pos[i+1][1] {
						pos[i+1][1]++
					} else if pos[i][1] < pos[i+1][1] {
						pos[i+1][1]--
					}
				}

				if pos[i][1] > pos[i+1][1]+1 || pos[i][1] < pos[i+1][1]-1 {
					if pos[i][1] > pos[i+1][1]+1 {
						pos[i+1][1]++
					} else if pos[i][1] < pos[i+1][1]-1 {
						pos[i+1][1]--
					}
					if pos[i][0] > pos[i+1][0] {
						pos[i+1][0]++
					} else if pos[i][0] < pos[i+1][0] {
						pos[i+1][0]--
					}
				}
			}

			visited[fmt.Sprintf("%d_%d", pos[len(pos)-1][0], pos[len(pos)-1][1])]++
		}
	}

	fmt.Println("visited:", len(visited)) // 2449
}
