package main

import (
	"log"
	"sort"
	"strconv"
)

func (d *Day1) part2() {

	const numTop = 3

	var cals []int
	for _, line := range d.InputLines {
		if line == "" || len(cals) == 0 {
			cals = append(cals, 0)
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			cals[len(cals)-1] += num
		}
	}

	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})

	sumCals := 0
	for i, cal := range cals {
		if i < numTop {
			sumCals += cal
		}
	}

	log.Println("sum calories:", sumCals) // 205805
}
