package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (d *Day7) part1() {

	var dir []string
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
		} else if !strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "dir") {
				sizes[strings.Join(dir, "/")] = 0
			} else {
				parts := strings.Split(line, " ")
				size, _ := strconv.Atoi(parts[0])
				sizes[strings.Join(dir, "/")] += size
			}
		}
	}

	for path := range sizes {
		for p, s := range sizes {
			if strings.HasPrefix(p, path) && path != p {
				sizes[path] += s
			}
		}
	}

	totalSize := 0
	for _, size := range sizes {
		if size <= 100000 {
			totalSize += size
		}
	}

	fmt.Println("total:", totalSize)
}
