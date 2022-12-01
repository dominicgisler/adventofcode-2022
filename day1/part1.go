package main

import (
	"log"
	"strconv"
)

func (d *Day1) part1() {

	calTemp := 0
	mostCal := 0
	for _, line := range d.InputLines {
		if line == "" {
			if calTemp > mostCal {
				mostCal = calTemp
			}
			calTemp = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calTemp += i
		}
	}

	log.Println("most calories:", mostCal) // 70613
}
