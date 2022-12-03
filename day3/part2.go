package main

import (
	"log"
)

func (d *Day3) part2() {

	var prio = 0
	for i, line := range d.InputLines {
		var badge rune
		if i%3 == 0 {
			set1 := line
			set2 := []rune(d.InputLines[i+1])
			set3 := []rune(d.InputLines[i+2])

			for _, r := range set1 {
				if d.contains(set2, r) && d.contains(set3, r) {
					badge = r
				}
			}
			prio += d.getPrio(badge)
		}
	}

	log.Println("total prio:", prio) // 2479
}
