package main

import (
	"aoc-2020/internal/filereader"
	"aoc-2020/internal/transformer"
	"fmt"
	"sort"
)

func main() {
	Day10()
}

func Day10() {
	fmt.Println("==============  Day10 =============")
	fmt.Println("==============  TEST  ============")

	testData, _error := transformer.SliceAtoi(filereader.ReadFile("./input/day-10/test.txt"))
	testData = append(testData, 0, maxInt(testData)+3)
	oneDiff, threeDiff := getDifferenceInJoltForAdapters(testData)
	output := oneDiff * threeDiff

	if output == 35 && _error == nil {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	testData2, _error := transformer.SliceAtoi(filereader.ReadFile("./input/day-10/test2.txt"))
	testData2 = append(testData2, 0, maxInt(testData2)+3)
	oneDiff, threeDiff = getDifferenceInJoltForAdapters(testData2)
	output = oneDiff * threeDiff

	if output == 220 && _error == nil {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	differentArrangements := countDifferentArrangements(testData)
	if differentArrangements == 8 && _error == nil {
		fmt.Println("Test succeeded:", differentArrangements)
	} else {
		fmt.Println("Test failed:", differentArrangements)
	}

	differentArrangements = countDifferentArrangements(testData2)
	if differentArrangements == 19208 && _error == nil {
		fmt.Println("Test succeeded:", differentArrangements)
	} else {
		fmt.Println("Test failed:", differentArrangements)
	}

	fmt.Println("============== OUTPUT ============")
	data, _err := transformer.SliceAtoi(filereader.ReadFile("./input/day-10/input.txt"))
	data = append(data, 0, maxInt(data)+3)
	if _err == nil {
		oneDiff, threeDiff = getDifferenceInJoltForAdapters(data)
		output = oneDiff * threeDiff
		fmt.Println("The multiplication of differences is:", output)

		fmt.Println("Different arrangements:", countDifferentArrangements(data))
	} else {
		fmt.Println(_err)
	}
}

func getDifferenceInJoltForAdapters(adapters []int) (int, int) {
	sort.Ints(adapters)
	differences := make(map[int]int)
	for index := 1; index < len(adapters); index++ {
		differences[adapters[index]-adapters[index-1]]++
	}
	return differences[1], differences[3]
}

func countDifferentArrangements(adapters []int) int {
	memo := map[int]int{}

	var countPossibilities func(startFrom int) int
	countPossibilities = func(startFrom int) int {
		if value, exists := memo[startFrom]; exists {
			return value
		}

		subInts := adapters[startFrom:]

		if len(subInts) <= 1 {
			return 1
		}

		first := subInts[0]
		withinThree := findIdxs(subInts, func(i int) bool {
			return i > first && i <= first+3
		})

		count := 0
		for _, idx := range withinThree {
			count += countPossibilities(startFrom + idx)
		}
		memo[startFrom] = count
		return count
	}

	return countPossibilities(0)
}

func findIdxs(ints []int, match func(i int) bool) []int {
	idxs := []int{}
	for idx, i := range ints {
		if match(i) {
			idxs = append(idxs, idx)
		}
	}
	return idxs
}

func maxInt(ints []int) (max int) {
	max = ints[0]
	for _, i := range ints[1:] {
		if i > max {
			max = i
		}
	}
	return
}
