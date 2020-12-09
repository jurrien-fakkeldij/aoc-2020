package main

import (
	"aoc-2020/internal/filereader"
	"aoc-2020/internal/transformer"
	"fmt"
	"sort"
)

func main() {
	Day9()
}

func Day9() {
	fmt.Println("============== Day9 ==============")
	fmt.Println("============== TEST ==============")

	testData, _error := transformer.SliceAtoi(filereader.ReadFile("./input/day-9/test.txt"))

	output := getWeakness(5, testData)

	if output == 127 && _error == nil {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = breakEncryption(output, testData)

	if output == 62 && _error == nil {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT ============")
	data, _err := transformer.SliceAtoi(filereader.ReadFile("./input/day-9/input.txt"))
	if _err == nil {
		fmt.Println("The weaknumber is:", getWeakness(25, data))
	} else {
		fmt.Println(_err)
	}
	fmt.Println("The weakness is:", breakEncryption(getWeakness(25, data), data))
}

func getWeakness(preamble int, data []int) int {
	weakNumber := 0
	index := 0
	for index, weakNumber = range data[preamble:] {
		if !isSumInSlice(weakNumber, data[index:index+preamble]) {
			return weakNumber
		}
	}
	return weakNumber
}

func breakEncryption(weakNumber int, data []int) int {
	weakness := 0
	for amountOfDigits := 2; amountOfDigits < len(data); amountOfDigits++ {
		for index := 0; index < len(data)-amountOfDigits; index++ {
			sum := 0
			//fmt.Println("range:", data[index:index+amountOfDigits])
			for _, number := range data[index : index+amountOfDigits] {
				sum += number
			}

			if sum == weakNumber {
				//fmt.Println("Found something:", data[index:index+amountOfDigits])
				sortedSlice := data[index : index+amountOfDigits]
				sort.Ints(sortedSlice)
				fmt.Println("Found something:", sortedSlice)
				return sortedSlice[0] + sortedSlice[len(sortedSlice)-1]
			}
		}
	}

	return weakness
}

func isSumInSlice(sum int, data []int) bool {
	for _, number1 := range data {
		for _, number2 := range data {
			if number1 == number2 {
				continue
			} else {
				if sum == number1+number2 {
					return true
				}
			}
		}
	}

	//fmt.Println("None found for:", sum, "in", data)
	return false
}
