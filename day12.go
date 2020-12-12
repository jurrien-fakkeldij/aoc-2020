package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"math"
)

func main() {
	Day12()
}

func Day12() {
	fmt.Println("==============  Day11 =============")
	ship := &Ship{Vector{0, 0}, Vector{10, 1}}
	testInstructions := filereader.ReadFile("./input/day-12/test.txt")
	instructions := filereader.ReadFile("./input/day-12/input.txt")

	fmt.Println("==============  TEST  ============")
	for _, instruction := range testInstructions {
		ship.translateInstruction(instruction)
	}

	if math.Abs(float64(ship.pos.x))+math.Abs(float64(ship.pos.y)) == 286 {
		fmt.Println("Correct location:", ship.pos)
	} else {
		fmt.Println("Wrong location:", ship.pos)
	}

	fmt.Println("============== OUTPUT ============")
	ship = &Ship{Vector{0, 0}, Vector{10, 1}}

	for _, instruction := range instructions {
		ship.translateInstruction(instruction)
	}

	fmt.Println("Manhattan distance:", math.Abs(float64(ship.pos.x))+math.Abs(float64(ship.pos.y)))
}
