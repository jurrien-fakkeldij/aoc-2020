package main

import (
	"aoc-2020/internal/filereader"
	"aoc-2020/internal/transformer"
	"fmt"
	"strings"
)

func main() {
	Day15()
}

func Day15() {
	fmt.Println("==============  Day15 =============")
	//testInput, _ := transformer.SliceAtoi(strings.Split(filereader.ReadFile("./input/day-15/test.txt")[0], ","))
	fmt.Println("==============  TEST  ============")

	//output := keepCountingTillTurn(testInput, 2020)
	/*if output == 436 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = keepCountingTillTurn(testInput, 30000000)
	if output == 175594 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}*/

	fmt.Println("============== OUTPUT ============")
	input, _ := transformer.SliceAtoi(strings.Split(filereader.ReadFile("./input/day-15/input.txt")[0], ","))
	//fmt.Println("Number2020:", keepCountingTillTurn(input, 2020))
	fmt.Println("Number30000000:", keepCountingTillTurn(input, 30000000))
}

func keepCountingTillTurn(input []int, turn int) int {
	spokenNumbers := make(map[int]int)
	spokenNumber := 0

	for i := 1; i < turn; i++ {
		if i < len(input) {
			spokenNumber = input[i-1]
			spokenNumbers[spokenNumber] = i
			//fmt.Println("turn:", i, "number:", spokenNumber)
		} else {
			if i == len(input) {
				spokenNumber = input[i-1]
			}
			//fmt.Println("turn:", i, "number:", spokenNumber)
			if value, present := spokenNumbers[spokenNumber]; !present {
				spokenNumbers[spokenNumber] = i
				spokenNumber = 0
			} else {
				spokenNumbers[spokenNumber] = i
				spokenNumber = i - value
			}
		}
		//fmt.Println(spokenNumbers)
	}
	return spokenNumber
}

func keepCountingTillTurnOld(input []int, turn int) int {
	spokenNumbers := intSlice{}
	spokenNumber := 0

	//fmt.Println("input:", input)
	for i := 0; i < turn; i++ {
		if i < len(input) {
			spokenNumber = input[i]
		} else {
			if spokenNumbers[:i-1].lastpos(spokenNumber) == -1 {
				spokenNumber = 0
			} else {
				spokenNumber = i - 1 - spokenNumbers[:i-1].lastpos(spokenNumber)
			}
		}
		//fmt.Println("turn:", i+1, "number:", spokenNumber)
		spokenNumbers = append(spokenNumbers, spokenNumber)
	}

	return spokenNumber
}

type intSlice []int

func (slice intSlice) lastpos(value int) int {
	revSlice := make([]int, len(slice))
	copy(revSlice, slice)
	for p, v := range reverse(revSlice) {
		if v == value {
			//fmt.Println("Found slice:", revSlice, "lastpos:", len(revSlice)-1-p, "value:", value)
			return len(revSlice) - 1 - p
		}
	}
	//fmt.Println("Not Found slice:", revSlice, "lastpos:", -1, "value:", value)
	return -1
}

func reverse(slice intSlice) intSlice {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
