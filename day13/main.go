package main

import (
	"encoding/json"
	"github.com/dominicgisler/adventofcode-2022/core"
)

type Day13 struct {
	core.Core
}

func main() {
	d := Day13{}
	d.Init()
	d.part1()
	d.part2()
}

func (d *Day13) compareList(list1 []any, list2 []any) (bool, bool) {
	done := false
	order := false
	for j := range list1 {
		if done {
			continue
		}
		if len(list2) <= j {
			return true, false
		}
		l1, l1i := list1[j].(float64)
		l2, l2i := list2[j].(float64)
		if l1i && l2i {
			if l1 < l2 {
				done = true
				order = true
			} else if l1 > l2 {
				done = true
				order = false
			}
		} else if !l1i && !l2i {
			done, order = d.compareList(list1[j].([]any), list2[j].([]any))
		} else if !l1i && l2i {
			done, order = d.compareList(list1[j].([]any), []any{l2})
		} else if l1i && !l2i {
			done, order = d.compareList([]any{l1}, list2[j].([]any))
		}
	}
	if !done && len(list1) < len(list2) {
		order = true
		done = true
	}
	return done, order
}

func (d *Day13) parseList(str string) []any {
	var res []any
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		panic(err)
	}
	return res
}
