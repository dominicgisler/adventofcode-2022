package main

import (
	"fmt"
	"sort"
	"strings"
)

func (d *Day11) part1() {

	type Monkey struct {
		Items           []int
		OperationType   string
		OperationFactor string
		TestNum         int
		TestTrue        int
		TestFalse       int
		Inspections     int
	}

	var monkeys []Monkey

	i := -1
	for _, line := range d.InputLines {
		switch {
		case strings.HasPrefix(line, "Monkey"):
			monkeys = append(monkeys, Monkey{})
			i++
		case strings.HasPrefix(line, "  Starting items"):
			nums := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "  Starting items: ", ""), " ", ""), ",")
			for _, num := range nums {
				monkeys[i].Items = append(monkeys[i].Items, d.StrToInt(num))
			}
		case strings.HasPrefix(line, "  Operation"):
			op := strings.Split(strings.ReplaceAll(line, "  Operation: new = old ", ""), " ")
			monkeys[i].OperationType = op[0]
			monkeys[i].OperationFactor = op[1]
		case strings.HasPrefix(line, "  Test"):
			test := strings.ReplaceAll(line, "  Test: divisible by ", "")
			monkeys[i].TestNum = d.StrToInt(test)
		case strings.HasPrefix(line, "    If true"):
			next := strings.ReplaceAll(line, "    If true: throw to monkey ", "")
			monkeys[i].TestTrue = d.StrToInt(next)
		case strings.HasPrefix(line, "    If false"):
			next := strings.ReplaceAll(line, "    If false: throw to monkey ", "")
			monkeys[i].TestFalse = d.StrToInt(next)
		}
	}

	for r := 0; r < 20; r++ {
		for i := range monkeys {
			d.DebugStr("Monkey %d:\n", i)
			for j := range monkeys[i].Items {
				monkeys[i].Inspections++
				d.DebugStr("  Monkey inspects an item with a worry level of %d.\n", monkeys[i].Items[j])
				fact := d.StrToInt(monkeys[i].OperationFactor)
				if monkeys[i].OperationFactor == "old" {
					fact = monkeys[i].Items[j]
				}
				switch monkeys[i].OperationType {
				case "+":
					monkeys[i].Items[j] += fact
					d.DebugStr("    Worry level increases by %d to %d.\n", fact, monkeys[i].Items[j])
				case "*":
					monkeys[i].Items[j] *= fact
					d.DebugStr("    Worry level is multiplied by %d to %d.\n", fact, monkeys[i].Items[j])
				}
				monkeys[i].Items[j] /= 3
				d.DebugStr("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", monkeys[i].Items[j])

				if monkeys[i].Items[j]%monkeys[i].TestNum == 0 {
					monkeys[monkeys[i].TestTrue].Items = append(monkeys[monkeys[i].TestTrue].Items, monkeys[i].Items[j])
					d.DebugStr("    Current worry level is divisible by %d.\n    Item with worry level %d is thrown to monkey %d.\n", monkeys[i].TestNum, monkeys[i].Items[j], monkeys[i].TestTrue)
				} else {
					monkeys[monkeys[i].TestFalse].Items = append(monkeys[monkeys[i].TestFalse].Items, monkeys[i].Items[j])
					d.DebugStr("    Current worry level is not divisible by %d.\n    Item with worry level %d is thrown to monkey %d.\n", monkeys[i].TestNum, monkeys[i].Items[j], monkeys[i].TestFalse)
				}
			}
			monkeys[i].Items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	fmt.Println("monkey business:", monkeys[0].Inspections*monkeys[1].Inspections) // 102399
}
