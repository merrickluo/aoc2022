package main

import (
	"strings"
	"testing"
)

const testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPart1(t *testing.T) {
	p1 := part1(strings.Split(testInput, "\n"))
	if p1 != 95437 {
		t.Errorf("expected p1: 95437, got: %d", p1)
	}
}

func TestPart2(t *testing.T) {
	p2 := part2(strings.Split(testInput, "\n"))
	if p2 != 24933642 {
		t.Errorf("expected p2: 24933642, got: %d", p2)
	}
}
