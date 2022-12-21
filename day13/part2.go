package main

import (
	"encoding/json"
	"fmt"
)

func (d *Day13) part2() {

	d.InputLines = append(d.InputLines, "", "[[2]]", "[[6]]")

	var newOrder []any
	for i, line := range d.InputLines {
		if i%3 == 0 {
			left := d.parseList(line)
			right := d.parseList(d.InputLines[i+1])
			newOrder = append(newOrder, left, right)
		}
	}

	hasChange := true
	for hasChange {
		hasChange = false
		for i := range newOrder {
			if i == len(newOrder)-1 || hasChange {
				continue
			}
			left := newOrder[i].([]any)
			right := newOrder[i+1].([]any)
			_, order := d.compareList(left, right)
			if order {
			} else {
				newOrder[i] = right
				newOrder[i+1] = left
				hasChange = true
			}
		}
	}

	var decoder = 1
	for i, line := range newOrder {
		jstr, _ := json.Marshal(line)
		if string(jstr) == "[[2]]" || string(jstr) == "[[6]]" {
			decoder *= i + 1
		}
	}

	fmt.Println("decoder key:", decoder) // 23958
}
