package main

import (
	"log"
	"strconv"
	"strings"
)

func (d *Day5) part1() {

	stacks := map[int][]string{}
	header := true
	for _, line := range d.InputLines {
		if line == "" {
			header = false
			continue
		}

		if header {
			for j, l := range line {
				if (j-1)%4 == 0 && strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(l)) {
					stacks[j/4+1] = append([]string{string(l)}, stacks[j/4+1]...)
				}
			}
			continue
		}

		line = strings.Replace(line, "move ", "", 1)
		line = strings.Replace(line, "from ", "", 1)
		line = strings.Replace(line, "to ", "", 1)
		nums := strings.Split(line, " ")

		cnt, _ := strconv.Atoi(nums[0])
		from, _ := strconv.Atoi(nums[1])
		to, _ := strconv.Atoi(nums[2])

		for i := 0; i < cnt; i++ {
			idx := len(stacks[from]) - 1
			el := stacks[from][idx]
			stacks[to] = append(stacks[to], el)
			stacks[from] = append(stacks[from][:idx], stacks[from][idx+1:]...)
		}
	}

	sol := ""
	for i := 1; i <= len(stacks); i++ {
		sol += stacks[i][len(stacks[i])-1]
	}

	log.Println("solution:", sol) // VJSFHWGFT
}
