package main

import (
	"fmt"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	lines := common.ReadLines("day8")
	trees := toTrees(lines)

	fmt.Println(part1(trees))
	fmt.Println(part2(trees))
}

func toTrees(lines []string) [][]int {
	trees := make([][]int, len(lines))

	for i, line := range lines {
		trees[i] = make([]int, len(line))
		for j, c := range line {
			trees[i][j] = int(c - '0')
		}
	}

	return trees
}

// visible tree count
func part1(trees [][]int) int {
	rows, cols := len(trees), len(trees[0])

	visible := 0
	for i := range trees {
		for j := range trees[i] {
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				visible += 1
				continue
			}

			// left
			k := j - 1
			for ; k >= 0; k-- {
				if trees[i][k] >= trees[i][j] {
					break
				}
			}
			if k < 0 {
				visible += 1
				continue
			}

			// right
			k = j + 1
			for ; k < cols; k++ {
				if trees[i][k] >= trees[i][j] {
					break
				}
			}
			if k >= cols {
				visible += 1
				continue
			}

			// top
			k = i - 1
			for ; k >= 0; k-- {
				if trees[k][j] >= trees[i][j] {
					break
				}
			}
			if k < 0 {
				visible += 1
				continue
			}

			// bottom
			k = i + 1
			for ; k < rows; k++ {
				if trees[k][j] >= trees[i][j] {
					break
				}
			}
			if k >= rows {
				visible += 1
				continue
			}
		}
	}

	return visible
}

// scenic score
func part2(trees [][]int) int {
	rows, cols := len(trees), len(trees[0])

	bestscore := 0
	for i := range trees {
		for j := range trees[i] {
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				continue // score is 0
			}

			// left
			treesLeft := 0
			for k := j - 1; k >= 0; k-- {
				treesLeft += 1
				if trees[i][k] >= trees[i][j] {
					break
				}
			}

			// right
			treesRight := 0
			for k := j + 1; k < cols; k++ {
				treesRight += 1
				if trees[i][k] >= trees[i][j] {
					break
				}
			}

			// top
			treesTop := 0
			for k := i - 1; k >= 0; k-- {
				treesTop += 1
				if trees[k][j] >= trees[i][j] {
					break
				}
			}

			// bottom
			treesBottom := 0
			for k := i + 1; k < rows; k++ {
				treesBottom += 1
				if trees[k][j] >= trees[i][j] {
					break
				}
			}

			score := treesLeft * treesRight * treesTop * treesBottom
			if score > bestscore {
				bestscore = score
			}
		}
	}

	return bestscore
}
