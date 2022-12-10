package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	lines := common.ReadLines("day4")
	pairs := toPairs(lines)

	fmt.Println(part1(pairs))
	fmt.Println(part2(pairs))
}

func toPairs(lines []string) []*common.Pair[*common.Pair[int]] {
	pairs := make([]*common.Pair[*common.Pair[int]], len(lines))

	for i, line := range lines {
		items := strings.Split(line, ",")
		pairs[i] = common.NewPair(toPair(items[0]), toPair(items[1]))
	}

	return pairs
}

func toPair(str string) *common.Pair[int] {
	items := strings.Split(str, "-")
	left, _ := strconv.Atoi(items[0])
	right, _ := strconv.Atoi(items[1])

	return common.NewPair(left, right)
}

func part1(pairs []*common.Pair[*common.Pair[int]]) int {
	counter := 0

	for _, pair := range pairs {
		if pair.Left.Left <= pair.Right.Left && pair.Left.Right >= pair.Right.Right {
			counter += 1
			continue
		}

		if pair.Left.Left >= pair.Right.Left && pair.Left.Right <= pair.Right.Right {
			counter += 1
			continue
		}
	}

	return counter
}

func part2(pairs []*common.Pair[*common.Pair[int]]) int {
	counter := 0

	for _, pair := range pairs {
		if pair.Left.Left <= pair.Right.Left && pair.Left.Right >= pair.Right.Left {
			counter += 1
			continue
		}

		if pair.Left.Left >= pair.Right.Left && pair.Right.Right >= pair.Left.Left {
			counter += 1
			continue
		}
	}

	return counter
}
