package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	p1 := part1(rucksacks)
	if p1 != 157 {
		t.Errorf("expect p1: 157: got: %v", p1)
	}
}

func TestPart2(t *testing.T) {
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	p2 := part2(rucksacks)
	if p2 != 70 {
		t.Errorf("expect p1: 70: got: %v", p2)
	}
}
