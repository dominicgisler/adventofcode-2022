package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (d *Day10) part1() {

	x := 1
	cycle := 0

	var stack []string
	for _, line := range d.InputLines {
		if line == "noop" {
			stack = append(stack, line)
		} else if strings.HasPrefix(line, "addx") {
			stack = append(stack, "noop", "noop", line)
		}
	}

	strength := 0
	checkCycle := 20
	for _, cmd := range stack {
		if cmd == "noop" {
			cycle++
		}

		if cycle == checkCycle {
			checkCycle += 40
			strength += cycle * x
		}

		if strings.HasPrefix(cmd, "addx") {
			parts := strings.Split(cmd, " ")
			num, _ := strconv.Atoi(parts[1])
			x += num
		}
	}

	fmt.Println("total strength:", strength) // 17020
}
