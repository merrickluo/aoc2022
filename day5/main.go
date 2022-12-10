package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	stacks := [][]string{
		{"P", "F", "M", "Q", "W", "G", "R", "T"},
		{"H", "F", "R"},
		{"P", "Z", "R", "V", "G", "H", "S", "D"},
		{"Q", "H", "P", "B", "F", "W", "G"},
		{"P", "S", "M", "J", "H"},
		{"M", "Z", "T", "H", "S", "R", "P", "L"},
		{"P", "T", "H", "N", "M", "L"},
		{"F", "D", "Q", "R"},
		{"D", "S", "C", "N", "L", "P", "H"},
	}

	moves := parseMoves(common.ReadLines("day5"))

	// fmt.Println(part1(stacks, moves))
	fmt.Println(part2(stacks, moves))
}

type Move struct {
	From  int
	To    int
	Count int
}

func (m *Move) String() string {
	return fmt.Sprintf("move %d from %d to %d,", m.Count, m.From, m.To)
}

func NewMove(from, to, count string) *Move {
	f, _ := strconv.Atoi(from)
	t, _ := strconv.Atoi(to)
	c, _ := strconv.Atoi(count)

	return &Move{
		From:  f,
		To:    t,
		Count: c,
	}
}

func parseMoves(lines []string) []*Move {
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	moves := make([]*Move, 0, len(lines))
	for _, line := range lines {
		res := r.FindStringSubmatch(line)
		if len(res) > 0 {
			moves = append(moves, NewMove(res[2], res[3], res[1]))
		}
	}

	return moves
}

func stacksTop(stacks [][]string) string {

	buf := ""
	for _, stack := range stacks {
		buf += stack[len(stack)-1]
	}

	return buf
}

func part1(stacks [][]string, moves []*Move) string {
	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			n := len(stacks[move.From-1]) - 1

			stacks[move.To-1] = append(stacks[move.To-1], stacks[move.From-1][n]) // push
			stacks[move.From-1] = stacks[move.From-1][:n]                         // pop
		}
	}

	return stacksTop(stacks)
}

func part2(stacks [][]string, moves []*Move) string {
	for _, move := range moves {
		n := len(stacks[move.From-1]) - move.Count
		stacks[move.To-1] = append(stacks[move.To-1], stacks[move.From-1][n:]...) // push
		stacks[move.From-1] = stacks[move.From-1][:n]                             // pop
	}

	return stacksTop(stacks)
}
