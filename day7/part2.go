package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (d *Day7) part2() {

	var dir []string
	var curr string
	sizes := map[string]int{"": 0}
	for _, line := range d.InputLines {
		if strings.HasPrefix(line, "$ cd") {
			parts := strings.Split(line, " ")
			path := parts[len(parts)-1]
			switch path {
			case "/":
				dir = []string{""}
			case "..":
				dir = append(dir[:len(dir)-1], dir[len(dir):]...)
			default:
				dir = append(dir, path)
			}
			curr = strings.Join(dir, "/")
		} else if !strings.HasPrefix(line, "$") && !strings.HasPrefix(line, "dir") {
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			sizes[curr] += size
			for p := range sizes {
				if strings.HasPrefix(curr, p) && curr != p {
					sizes[p] += size
				}
			}
		}
	}

	total := 70000000
	needed := 30000000
	unused := total - sizes[""]
	diff := needed - unused

	found := needed
	for _, size := range sizes {
		if size >= diff && size < found {
			found = size
		}
	}

	fmt.Println("size to clean:", found)
}
