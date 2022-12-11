package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	lines := common.ReadLines("day10")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func eval(lines []string) []int {
	x, history := 1, []int{1}

	for _, line := range lines {
		items := strings.Split(line, " ")
		switch items[0] {
		case "noop":
			history = append(history, x)
		case "addx":
			v, _ := strconv.Atoi(items[1])
			history = append(history, x, x+v)
			x = x + v
		}
	}

	return history
}

func part1(lines []string) int {
	history := eval(lines)

	sum := 0
	for i := 20; i <= 220; i += 40 {
		sum += history[i-1] * i
	}

	return sum
}

func part2(lines []string) string {
	history := eval(lines)
	buf := bytes.Buffer{}

	rows := len(history) / 40
	for i := 0; i < rows; i++ {
		for j := 0; j < 40; j++ {
			x := history[i*40+j]
			if math.Abs(float64(x-j)) < 2 {
				buf.WriteString("#")
				continue
			}
			buf.WriteString(".")
		}
		buf.WriteString("\n")
	}

	return buf.String()
}
