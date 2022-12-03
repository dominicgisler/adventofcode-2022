package main

import (
	"log"
)

func (d *Day3) part1() {

	var prio = 0
	for _, line := range d.InputLines {
		var wrong rune
		var comp1 []rune
		var comp2 []rune

		for i, r := range line {
			if i < len(line)/2 {
				comp1 = append(comp1, r)
			} else {
				comp2 = append(comp2, r)
			}
		}

		for _, r := range comp1 {
			if d.contains(comp2, r) {
				wrong = r
			}
		}

		prio += d.getPrio(wrong)
	}

	log.Println("total prio:", prio) // 7875
}
