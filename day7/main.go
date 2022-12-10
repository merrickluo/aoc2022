package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/merrickluo/aoc2022/common"
)

func main() {
	outputs := common.ReadLines("day7")

	fmt.Println(part1(outputs))
	fmt.Println(part2(outputs))
}

type Directory struct {
	name string
	size int
}

func buildFileTree(output []string) *common.Tree[Directory] {
	root := &common.Tree[Directory]{
		Value: Directory{
			name: "/",
			size: 0,
		},
	}

	pwd := root
	dirstack := []*common.Tree[Directory]{}

	for _, line := range output {
		items := strings.Split(line, " ")
		switch items[0] {
		case "$":
			if items[1] != "cd" { // ignores "ls"
				continue
			}

			switch items[2] {
			case "..":
				n := len(dirstack) - 1
				pwd = dirstack[n]
				dirstack = dirstack[:n]
			case "/":
				dirstack = []*common.Tree[Directory]{root}
				pwd = root
			default:
				// don't care about edge case
				for _, dir := range pwd.Children {
					if dir.Value.name == items[2] {
						dirstack = append(dirstack, pwd)
						pwd = dir
						break
					}
				}
			}
		case "dir":
			dir := Directory{name: items[1], size: 0}
			subtree := &common.Tree[Directory]{Value: dir}
			pwd.Children = append(pwd.Children, subtree)
		default:
			size, _ := strconv.Atoi(items[0])
			pwd.Value.size += size
		}
	}

	return root
}

func treeSize(root *common.Tree[Directory]) int {
	totalSize := root.Value.size

	if root.Children != nil {
		for _, child := range root.Children {
			totalSize += treeSize(child)
		}
	}

	return totalSize
}

func part1(output []string) int {
	root := buildFileTree(output)

	totalSize := 0
	root.Walk(func(t *common.Tree[Directory]) {
		size := treeSize(t)
		if size < 100000 {
			totalSize += size
		}
	})

	return totalSize
}

func part2(output []string) int {
	root := buildFileTree(output)

	sizes := []int{}
	used := treeSize(root)
	root.Walk(func(t *common.Tree[Directory]) {
		sizes = append(sizes, treeSize(t))
	})

	sort.Ints(sizes)
	for _, size := range sizes {
		if 70000000-used+size > 30000000 {
			return size
		}
	}

	return 0
}
