package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day2/day2.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	pairs := [][]string{}
	for _, line := range lines {
		pairs = append(pairs, strings.Split(line, " "))
	}

	fmt.Println(part1(pairs))
	fmt.Println(part2(pairs))
}

func part1(pairs [][]string) int {
	totalScore := 0
	for _, pair := range pairs {
		totalScore += score1(pair[0], pair[1])
	}

	return totalScore
}

func part2(pairs [][]string) int {
	totalScore := 0
	for _, pair := range pairs {
		totalScore += score2(pair[0], pair[1])
	}

	return totalScore
}

func score2(op, end string) int {
	switch end {
	case "X":
		switch op {
		case "A":
			return 3
		case "B":
			return 1
		case "C":
			return 2
		}
	case "Y":
		switch op {
		case "A":
			return 4
		case "B":
			return 5
		case "C":
			return 6
		}
	case "Z":
		switch op {
		case "A":
			return 8
		case "B":
			return 9
		case "C":
			return 7
		}
	}

	return 0
}

func score1(op, you string) int {
	if sameShape(op, you) { // draw
		return 3 + shapeScore(you)
	}

	if roundWin(op, you) {
		return 6 + shapeScore(you)
	}

	return shapeScore(you)
}

func roundWin(op, you string) bool {
	return op == "A" && you == "Y" ||
		op == "B" && you == "Z" ||
		op == "C" && you == "X"
}

func sameShape(op, you string) bool {
	return op == "A" && you == "X" ||
		op == "B" && you == "Y" ||
		op == "C" && you == "Z"
}

func shapeScore(shape string) int {
	switch shape {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	return 0
}
