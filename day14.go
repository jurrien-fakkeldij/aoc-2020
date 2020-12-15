package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
)

func main() {
	Day14()
}

func Day14() {
	fmt.Println("==============  Day14 =============")
	testInput := filereader.ReadFile("./input/day-14/test.txt")
	testInput2 := filereader.ReadFile("./input/day-14/test2.txt")
	fmt.Println("==============  TEST  ============")
	ship := &Ship{}
	output := ship.initializeDockProgram(testInput)
	if output == 165 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = ship.initializeDockProgramv2(testInput2)
	if output == 208 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT ============")
	input := filereader.ReadFile("./input/day-14/input.txt")
	fmt.Println("Values left:", ship.initializeDockProgram(input))
	fmt.Println("Values left:", ship.initializeDockProgramv2(input))
}
