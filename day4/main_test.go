package main

import "testing"

func TestPart1(t *testing.T) {
	pairs := toPairs([]string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	})

	p1 := part1(pairs)
	if p1 != 2 {
		t.Errorf("expected p1: 2, got: %d", p1)
	}
}

func TestPart2(t *testing.T) {
	pairs := toPairs([]string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	})

	p2 := part2(pairs)
	if p2 != 4 {
		t.Errorf("expected p2: 4, got: %d", p2)
	}
}
