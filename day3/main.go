package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day3/day3.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(rucksacks []string) int {
	priorities := 0

	for _, rucksack := range rucksacks {
		size := len(rucksack)
		com1, com2 := rucksack[:size/2], rucksack[size/2:]
	ruck:
		for _, c1 := range com1 {
			for _, c2 := range com2 {
				if c1 == c2 { // only one
					priorities += priority(c1)
					break ruck
				}
			}
		}
	}

	return priorities
}

func priority(c rune) int {
	if c < 'a' {
		return int(c-'A') + 27
	}

	return int(c-'a') + 1
}

func part2(rucksacks []string) int {
	priorities := 0

	for i := 0; i < len(rucksacks); i += 3 {

	ruck:
		for _, c1 := range rucksacks[i] {
			for _, c2 := range rucksacks[i+1] {
				for _, c3 := range rucksacks[i+2] {
					if c1 == c2 && c2 == c3 {
						priorities += priority(c2)
						break ruck
					}
				}
			}
		}
	}

	return priorities
}
