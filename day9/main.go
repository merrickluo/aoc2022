package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	moves := toMoves(common.ReadLines("day9"))

	fmt.Println(part1(moves))
	fmt.Println(part2(moves))
}

type Move struct {
	direction string
	steps     int
}

func toMoves(lines []string) []Move {
	moves := make([]Move, len(lines))

	for i, line := range lines {
		items := strings.Split(line, " ")
		moves[i].direction = items[0]

		c, _ := strconv.Atoi(items[1])
		moves[i].steps = c
	}

	return moves
}

func part1(moves []Move) int {
	head, tail := common.NewPair(0, 0), common.NewPair(0, 0)

	visited := map[string]struct{}{}
	for _, move := range moves {
		for i := 0; i < move.steps; i++ {
			switch move.direction {
			case "R":
				head.Left++
			case "U":
				head.Right++
			case "L":
				head.Left--
			case "D":
				head.Right--
			}

			if !adj(head, tail) {
				visited[tail.String()] = struct{}{}

				if head.Left == tail.Left {
					tail.Right = moveTowards(tail.Right, head.Right)
				} else if head.Right == tail.Right {
					tail.Left = moveTowards(tail.Left, head.Left)
				} else {
					tail.Left = moveTowards(tail.Left, head.Left)
					tail.Right = moveTowards(tail.Right, head.Right)
				}
			}
		}
	}

	return len(visited) + 1
}

func moveTowards(pos, towards int) int {
	if pos < towards {
		return pos + 1
	} else {
		return pos - 1
	}
}

func adj(head, tail *common.Pair[int]) bool {
	return math.Abs(float64(head.Left-tail.Left)) <= 1 &&
		math.Abs(float64(head.Right-tail.Right)) <= 1
}

func part2(moves []Move) int {
	visited := map[string]struct{}{}
	knobs := make([]*common.Pair[int], 10)

	for i := range knobs {
		knobs[i] = common.NewPair(0, 0)
	}

	var head *common.Pair[int]
	for _, move := range moves {
		for i := 0; i < move.steps; i++ {
			head = knobs[0]

			switch move.direction {
			case "R":
				head.Left++
			case "U":
				head.Right++
			case "L":
				head.Left--
			case "D":
				head.Right--
			}

			for i := 1; i < 10; i++ {
				knob := knobs[i]
				if !adj(head, knob) {
					if head.Left == knob.Left {
						knob.Right = moveTowards(knob.Right, head.Right)
					} else if head.Right == knob.Right {
						knob.Left = moveTowards(knob.Left, head.Left)
					} else {
						knob.Left = moveTowards(knob.Left, head.Left)
						knob.Right = moveTowards(knob.Right, head.Right)
					}
				}
				head = knob
			}
			// head is now tail
			visited[head.String()] = struct{}{}
		}
	}

	return len(visited)
}
