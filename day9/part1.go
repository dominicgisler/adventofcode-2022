package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (d *Day9) part1() {

	headPos := []int{0, 0}
	tailPos := []int{0, 0}

	visited := map[string]int{}
	for _, line := range d.InputLines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		num, _ := strconv.Atoi(parts[1])

		for i := 0; i < num; i++ {
			switch dir {
			case "R":
				headPos[0]++
			case "L":
				headPos[0]--
			case "U":
				headPos[1]++
			case "D":
				headPos[1]--
			}

			if headPos[0] > tailPos[0]+1 || headPos[0] < tailPos[0]-1 {
				if headPos[0] > tailPos[0]+1 {
					tailPos[0]++
				} else if headPos[0] < tailPos[0]-1 {
					tailPos[0]--
				}
				if headPos[1] > tailPos[1] {
					tailPos[1]++
				} else if headPos[1] < tailPos[1] {
					tailPos[1]--
				}
			}

			if headPos[1] > tailPos[1]+1 || headPos[1] < tailPos[1]-1 {
				if headPos[1] > tailPos[1]+1 {
					tailPos[1]++
				} else if headPos[1] < tailPos[1]-1 {
					tailPos[1]--
				}
				if headPos[0] > tailPos[0] {
					tailPos[0]++
				} else if headPos[0] < tailPos[0] {
					tailPos[0]--
				}
			}

			visited[fmt.Sprintf("%d_%d", tailPos[0], tailPos[1])]++
		}
	}

	fmt.Println("visited:", len(visited)) // 6236
}
