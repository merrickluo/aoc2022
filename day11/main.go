package main

import (
	"fmt"
	"sort"
)

func main() {
	monkeys := []*Monkey{
		{
			items: []int64{80},
			operation: func(old int64) int64 {
				return old * 5
			},
			trueTarget:  4,
			falseTarget: 3,
			divider:     2,
		},
		{
			items: []int64{75, 83, 74},
			operation: func(old int64) int64 {
				return old + 7
			},
			trueTarget:  5,
			falseTarget: 6,
			divider:     7,
		},
		{
			items: []int64{86, 67, 61, 96, 52, 63, 73},
			operation: func(old int64) int64 {
				return old + 5
			},
			trueTarget:  7,
			falseTarget: 0,
			divider:     3,
		},
		{
			items: []int64{85, 83, 55, 85, 57, 70, 85, 52},
			operation: func(old int64) int64 {
				return old + 8
			},
			trueTarget:  1,
			falseTarget: 5,
			divider:     17,
		},
		{
			items: []int64{67, 75, 91, 72, 89},
			operation: func(old int64) int64 {
				return old + 4
			},
			trueTarget:  3,
			falseTarget: 1,
			divider:     11,
		},
		{
			items: []int64{66, 64, 68, 92, 68, 77},
			operation: func(old int64) int64 {
				return old * 2
			},
			trueTarget:  6,
			falseTarget: 2,
			divider:     19,
		},
		{
			items: []int64{97, 94, 79, 88},
			operation: func(old int64) int64 {
				return old * old
			},
			trueTarget:  2,
			falseTarget: 7,
			divider:     5,
		},
		{
			items: []int64{77, 85},
			operation: func(old int64) int64 {
				return old + 6
			},
			trueTarget:  4,
			falseTarget: 0,
			divider:     13,
		},
	}

	fmt.Println(part1(monkeys))
	fmt.Println(part2(monkeys))
}

type Monkey struct {
	items       []int64
	operation   func(old int64) int64
	inspected   int64
	divider     int64
	trueTarget  int
	falseTarget int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("%v", m.items)
}

func (m *Monkey) Play(monkeys []*Monkey, manage func(lvl int64) int64) {
	for _, item := range m.items {
		m.inspected++

		lvl := m.operation(item)
		if manage != nil {
			lvl = manage(lvl)
		}

		target := m.falseTarget
		if lvl%m.divider == 0 {
			target = m.trueTarget
		}

		monkeys[target].items = append(monkeys[target].items, lvl)
	}
	m.items = []int64{}
}

func part1(monkeys []*Monkey) int64 {
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.Play(monkeys, func(lvl int64) int64 { return lvl / 3 })
		}
	}

	count := len(monkeys)
	times := make([]int64, count)
	for i, m := range monkeys {
		times[i] = m.inspected
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})

	return times[count-2] * times[count-1]
}

func part2(monkeys []*Monkey) int64 {
	modder := int64(1)
	for _, m := range monkeys {
		modder *= m.divider
	}

	for i := 0; i < 1000; i++ {
		for _, m := range monkeys {
			m.Play(monkeys, func(lvl int64) int64 {
				return lvl % modder
			})
		}
	}

	count := len(monkeys)
	times := make([]int64, count)
	for i, m := range monkeys {
		times[i] = m.inspected
	}

	fmt.Println(times)
	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})

	return times[count-2] * times[count-1]
}
