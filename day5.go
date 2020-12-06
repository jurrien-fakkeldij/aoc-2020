package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Day5()
}

func Day5() {
	fmt.Println("============== Day5 ==============")
	fmt.Println("============== TEST ==============")
	testBoardingPasses := filereader.ReadFile("./input/day-5/test.txt")
	highest := findHighestSeatID(testBoardingPasses)

	if highest == 820 {
		fmt.Println("Test succeeded:", highest)
	} else {
		fmt.Println("Test failed:", highest)
	}

	fmt.Println("============== OUTPUT ============")
	actualBoardingPasses := filereader.ReadFile("./input/day-5/input.txt")
	highest = findHighestSeatID(actualBoardingPasses)
	fmt.Println("Highest Seat:", highest)

	allSeats := findAllSeatIDs(actualBoardingPasses)
	sort.Ints(allSeats)
	foundSeatID := 0
	for index, seatID := range allSeats {
		if index+1 == len(allSeats) {
			//noop - fmt.Println("End Of Line")
		} else if seatID+1 != allSeats[index+1] {
			foundSeatID = seatID + 1
		}
	}

	fmt.Println("Found Seat:", foundSeatID)
}

func findAllSeatIDs(boardingPasses []string) []int {
	var seatIDS []int
	seatIDS = make([]int, len(boardingPasses))
	for index, boardingPass := range boardingPasses {
		seatIDS[index] = findSeatID(boardingPass)
	}
	return seatIDS
}

func findHighestSeatID(boardingPasses []string) int {
	highest := 0
	for _, boardingPass := range boardingPasses {
		current := findSeatID(boardingPass)
		if highest < current {
			highest = current
		}
	}
	return highest
}

func findSeatID(boardingPass string) int {
	row, error := findRow(boardingPass[0:7])

	if error != nil {
		fmt.Println("Error finding row: ", error)
		return 0
	}

	column, error := findColumn(boardingPass[7:10])

	if error != nil {
		fmt.Println("Error finding column: ", error)
		return 0
	}

	seatID := (row * 8) + column

	return seatID
}

func findRow(rowIdentifier string) (int, error) {
	binRowIdentifier := strings.ReplaceAll(strings.ReplaceAll(rowIdentifier, "B", "1"), "F", "0")
	column, error := strconv.ParseInt(binRowIdentifier, 2, 8)
	return int(column), error
}

func findColumn(columnIdentifier string) (int, error) {
	binColumnIdentifier := strings.ReplaceAll(strings.ReplaceAll(columnIdentifier, "R", "1"), "L", "0")
	column, error := strconv.ParseInt(binColumnIdentifier, 2, 0)
	return int(column), error
}
