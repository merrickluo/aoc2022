package main

import (
	"fmt"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	lines := common.ReadLines("day6")

	fmt.Println(part1(lines[0], 4))
	fmt.Println(part1(lines[0], 14))
}

func part1(signal string, length int) int {
	cs := []rune{}
	for i, c := range signal {
		for j, cc := range cs {
			if c == cc {
				cs = cs[j+1:]
				break
			}
		}

		cs = append(cs, c)
		if len(cs) == length {
			return i + 1
		}
	}

	return 0
}
