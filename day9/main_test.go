package main

import (
	"strings"
	"testing"
)

const testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestPart1(t *testing.T) {
	moves := toMoves(strings.Split(testInput, "\n"))

	p1 := part1(moves)

	if p1 != 13 {
		t.Errorf("expected p1: 13, got %d", p1)
	}
}

const testInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestPart2(t *testing.T) {
	moves := toMoves(strings.Split(testInput2, "\n"))

	p2 := part2(moves)
	if p2 != 36 {
		t.Errorf("expected p2: 36, got %d", p2)
	}
}
