package main

import (
	"log"
	"strconv"
	"strings"
)

func (d *Day4) part2() {

	cnt := 0
	for _, line := range d.InputLines {
		s := strings.Split(line, ",")
		s1 := strings.Split(s[0], "-")
		s2 := strings.Split(s[1], "-")

		s1f, _ := strconv.Atoi(s1[0])
		s1t, _ := strconv.Atoi(s1[1])
		s2f, _ := strconv.Atoi(s2[0])
		s2t, _ := strconv.Atoi(s2[1])

		if s1t >= s2f && s1f <= s2f || s2t >= s1f && s2f <= s1f {
			cnt++
		}
	}

	log.Println("cnt:", cnt) // 952
}
