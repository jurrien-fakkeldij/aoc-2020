package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
)

func main() {
	Day3()
}

type Vector struct {
	x int
	y int
}

func Day3() {
	fmt.Println("============== Day3 ==============")
	fmt.Println("============== Part1 =============")
	testMap := filereader.ReadFile("./input/day-3/test.txt")
	actualMap := filereader.ReadFile("./input/day-3/input.txt")

	count := countHitTrees(testMap, Vector{3, 1})
	if count == 7 {
		fmt.Println("Test succeeded:", count)
	} else {
		fmt.Println("Test failed:", count)
	}

	count = countHitTrees(actualMap, Vector{3, 1})
	fmt.Println("Number of hit trees:", count)

	fmt.Println("============== Part2 =============")
	count = 1
	directions := []Vector{Vector{1, 1}, Vector{3, 1}, Vector{5, 1}, Vector{7, 1}, Vector{1, 2}}
	for _, direction := range directions {
		count *= countHitTrees(actualMap, direction)
	}
	fmt.Println("Number of hit trees:", count)
}

func countHitTrees(groundMap []string, direction Vector) int {
	count := 0
	pointer := 0
	for index, line := range groundMap {
		runeLine := []rune(line)
		if index == 0 {
			//noop
		} else if direction.y != 1 && index%direction.y == 1 {
			continue
		} else {
			if string(line[pointer]) == "#" {
				count++
			}
		}
		pointer += direction.x
		if pointer >= len(runeLine) {
			pointer = pointer % len(runeLine)
		}
	}

	return count
}
