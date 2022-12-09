package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day1/day1.txt")

	lines := strings.Split(string(data), "\n")
	calories := [][]int{{}}
	elf := 0

	for _, line := range lines {
		if line != "" {
			calorie, _ := strconv.Atoi(line)
			calories[elf] = append(calories[elf], calorie)
			continue
		}
		elf++
		calories = append(calories, []int{})
	}

	fmt.Println(part1(calories))
	fmt.Println(part2(calories))
}

func part1(calories [][]int) int {
	max := 0

	for _, items := range calories {
		sum := 0
		for _, item := range items {
			sum += item
		}
		if sum > max {
			max = sum
		}
	}

	return max
}

func part2(calories [][]int) int {
	sums := make([]int, len(calories))

	for i, items := range calories {
		for _, item := range items {
			sums[i] += item
		}
	}

	sort.Ints(sums)

	top3sum := 0
	for _, calorie := range sums[len(sums)-3:] {
		top3sum += calorie
	}

	return top3sum
}
