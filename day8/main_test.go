package main

import (
	"strings"
	"testing"
)

const testInput = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	trees := toTrees(strings.Split(testInput, "\n"))

	p1 := part1(trees)
	if p1 != 21 {
		t.Errorf("expected p1: 21, got: %d", p1)
	}
}

func TestPart2(t *testing.T) {
	trees := toTrees(strings.Split(testInput, "\n"))

	p2 := part2(trees)
	if p2 != 8 {
		t.Errorf("expected p2: 8, got: %d", p2)
	}
}
