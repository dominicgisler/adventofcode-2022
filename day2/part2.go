package main

import (
	"log"
	"strings"
)

func (d *Day2) part2() {

	scores := map[string]map[string]int{
		"X": {
			"A": 3 + 0,
			"B": 1 + 0,
			"C": 2 + 0,
		},
		"Y": {
			"A": 1 + 3,
			"B": 2 + 3,
			"C": 3 + 3,
		},
		"Z": {
			"A": 2 + 6,
			"B": 3 + 6,
			"C": 1 + 6,
		},
	}

	score := 0
	for _, line := range d.InputLines {
		plays := strings.Split(line, " ")
		score += scores[plays[1]][plays[0]]
	}

	log.Println("score:", score) // 12429
}
