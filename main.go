package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
)

func main() {
	text := filereader.ReadFile("input/test/test.txt")

	// and then a loop iterates through
	// and prints each of the slice values.
	for _, eachLn := range text {
		fmt.Println(eachLn)
	}
}
