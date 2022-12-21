package main

import (
	"fmt"
)

func (d *Day13) part1() {

	var rightOrder []int
	var idx = 0
	for i, line := range d.InputLines {
		if i%3 == 0 {
			idx++
			left := d.parseList(line)
			right := d.parseList(d.InputLines[i+1])
			_, order := d.compareList(left, right)
			if order {
				rightOrder = append(rightOrder, idx)
			}
		}
	}

	var sum int
	for _, idx := range rightOrder {
		sum += idx
	}
	fmt.Println("sum:", sum) // 5588
}
