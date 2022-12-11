package main

import "testing"

func TestPart1(t *testing.T) {
	monkeys := []*Monkey{
		{
			items: []int64{79, 98},
			operation: func(old int64) int64 {
				return old * 19
			},
			trueTarget:  2,
			falseTarget: 3,
			divider:     23,
		},
		{
			items: []int64{54, 65, 75, 74},
			operation: func(old int64) int64 {
				return old + 6
			},
			trueTarget:  2,
			falseTarget: 0,
			divider:     19,
		},
		{
			items: []int64{79, 60, 97},
			operation: func(old int64) int64 {
				return old * old
			},
			trueTarget:  1,
			falseTarget: 3,
			divider:     13,
		},
		{
			items: []int64{74},
			operation: func(old int64) int64 {
				return old + 3
			},
			trueTarget:  0,
			falseTarget: 1,
			divider:     17,
		},
	}

	p1 := part1(monkeys)
	if p1 != 10605 {
		t.Errorf("expected p1: 10605, got: %d", p1)
	}
}

func TestPart2(t *testing.T) {
	monkeys := []*Monkey{
		{
			items: []int64{79, 98},
			operation: func(old int64) int64 {
				return old * 19
			},
			trueTarget:  2,
			falseTarget: 3,
			divider:     23,
		},
		{
			items: []int64{54, 65, 75, 74},
			operation: func(old int64) int64 {
				return old + 6
			},
			trueTarget:  2,
			falseTarget: 0,
			divider:     19,
		},
		{
			items: []int64{79, 60, 97},
			operation: func(old int64) int64 {
				return old * old
			},
			trueTarget:  1,
			falseTarget: 3,
			divider:     13,
		},
		{
			items: []int64{74},
			operation: func(old int64) int64 {
				return old + 3
			},
			trueTarget:  0,
			falseTarget: 1,
			divider:     17,
		},
	}

	p2 := part2(monkeys)
	if p2 != 2713310158 {
		t.Errorf("expected p2: 2713310158, got: %d", p2)
	}
}
