package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (d *Day10) part2() {

	x := 1
	cycle := 0
	crt := ""

	var stack []string
	for _, line := range d.InputLines {
		if line == "noop" {
			stack = append(stack, line)
		} else if strings.HasPrefix(line, "addx") {
			stack = append(stack, "noop", "noop", line)
		}
	}

	for _, cmd := range stack {
		if cmd == "noop" {
			cycle++

			if x >= cycle%40-2 && x <= cycle%40 {
				crt += "#"
			} else {
				crt += "."
			}

			if cycle%40 == 0 {
				crt += "\n"
			}
		}

		if strings.HasPrefix(cmd, "addx") {
			parts := strings.Split(cmd, " ")
			num, _ := strconv.Atoi(parts[1])
			x += num
		}
	}

	fmt.Println(crt) // RLEZFLGE
}
