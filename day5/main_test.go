package main

import (
	"strings"
	"testing"
)

const testInput = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestPart1(t *testing.T) {
	stacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	moves := parseMoves(strings.Split(testInput, "\n"))

	p1 := part1(stacks, moves)

	if p1 != "CMZ" {
		t.Errorf("expected p1: CMZ, got: %v", p1)
	}
}

func TestPart2(t *testing.T) {
	stacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	moves := parseMoves(strings.Split(testInput, "\n"))

	p2 := part2(stacks, moves)

	if p2 != "MCD" {
		t.Errorf("expected p2: MCD, got: %s", p2)
	}
}
