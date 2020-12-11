package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strings"
)

type Vector struct {
	x int
	y int
}

func main() {
	Day11()
}

func Day11() {
	fmt.Println("==============  Day11 =============")
	testDataSeatingArea := filereader.ReadFile("./input/day-11/test.txt")
	dataSeatingArea := filereader.ReadFile("./input/day-11/input.txt")
	fmt.Println("==============  TEST  ============")
	seatingArea, size := createSeatingArea(testDataSeatingArea)

	_, numberOfTakenSeats := simulateSeating(seatingArea, size, -1, false, false)
	if numberOfTakenSeats == 37 {
		fmt.Println("Test succeeded:", numberOfTakenSeats)
	} else {
		fmt.Println("Test failed:", numberOfTakenSeats)
	}

	_, numberOfTakenSeats = simulateSeating(seatingArea, size, -1, false, true)
	if numberOfTakenSeats == 26 {
		fmt.Println("Test succeeded:", numberOfTakenSeats)
	} else {
		fmt.Println("Test failed:", numberOfTakenSeats)
	}
	fmt.Println("============== OUTPUT ============")
	seatingArea, size = createSeatingArea(dataSeatingArea)
	_, numberOfTakenSeats = simulateSeating(seatingArea, size, -1, false, false)
	fmt.Println("Number of takenSeats:", numberOfTakenSeats)
	_, numberOfTakenSeats = simulateSeating(seatingArea, size, -1, false, true)
	fmt.Println("Number of takenSeats based on Line of Sight:", numberOfTakenSeats)
}

func createSeatingArea(inputData []string) (map[Vector]rune, Vector) {
	seatingArea := make(map[Vector]rune)
	size := Vector{0, 0}
	for rowID, row := range inputData {
		for colID, column := range []rune(row) {
			seatingArea[Vector{colID, rowID}] = column
		}

		xSize := len([]rune(row))
		if xSize > size.x {
			size.x = xSize
		}
	}

	size.y = len(inputData)

	return seatingArea, size
}

func simulateSeating(seatingArea map[Vector]rune, size Vector, sims int, debug bool, losCheck bool) (int, int) {
	newSeatingArea := make(map[Vector]rune)
	changes := 0
	numberOfSeatsTaken := 0
	takenSeat := rune('#')
	empty := rune('.')

	for pos, char := range seatingArea {
		if char != empty {
			if !losCheck {
				newSeatingArea[pos] = checkNeighboursForAction(seatingArea, pos)
			} else {
				newSeatingArea[pos] = checkLOSForAction(seatingArea, pos)
			}
			if newSeatingArea[pos] != char {
				changes++
			}

			if newSeatingArea[pos] == takenSeat {
				numberOfSeatsTaken++
			}
		} else {
			newSeatingArea[pos] = char
		}
	}
	if debug {
		printSeatingArea(newSeatingArea, size)
	}
	if changes != 0 && sims != 0 {
		//fmt.Println("new sim -> changes", changes)
		changes, numberOfSeatsTaken = simulateSeating(newSeatingArea, size, sims, debug, losCheck)
	}

	return changes, numberOfSeatsTaken
}

func checkNeighboursForAction(seatingArea map[Vector]rune, pos Vector) rune {
	seat := rune('L')
	//empty := rune('.')
	takenSeat := rune('#')

	upLeftPos := Vector{pos.x - 1, pos.y - 1}
	upPos := Vector{pos.x, pos.y - 1}
	upRightPos := Vector{pos.x + 1, pos.y - 1}
	leftPos := Vector{pos.x - 1, pos.y}
	rightPos := Vector{pos.x + 1, pos.y}
	bottomLeftPos := Vector{pos.x - 1, pos.y + 1}
	bottomPos := Vector{pos.x, pos.y + 1}
	bottomRightPos := Vector{pos.x + 1, pos.y + 1}
	charactersAround := []rune{
		seatingArea[upLeftPos],
		seatingArea[upPos],
		seatingArea[upRightPos],
		seatingArea[leftPos],
		seatingArea[rightPos],
		seatingArea[bottomLeftPos],
		seatingArea[bottomPos],
		seatingArea[bottomRightPos]}

	if strings.Count(string(charactersAround), string(takenSeat)) >= 4 {
		return seat
	} else if seatingArea[pos] == seat && strings.Count(string(charactersAround), string(takenSeat)) >= 1 {
		return seat
	}

	return takenSeat
}

func checkLOSForAction(seatingArea map[Vector]rune, pos Vector) rune {
	upLeftPos := Vector{-1, -1}
	upPos := Vector{0, -1}
	upRightPos := Vector{1, -1}
	leftPos := Vector{-1, 0}
	rightPos := Vector{1, 0}
	bottomLeftPos := Vector{-1, 1}
	bottomPos := Vector{0, 1}
	bottomRightPos := Vector{1, 1}

	directions := []Vector{
		upLeftPos,
		upPos,
		upRightPos,
		leftPos,
		rightPos,
		bottomLeftPos,
		bottomPos,
		bottomRightPos}

	seat := rune('L')
	empty := rune('.')
	takenSeat := rune('#')

	takenSeats := 0
	seats := 0
	attempt := 1
	debug := false
	for _, direction := range directions {
		attempt = 1
		if debug {
			fmt.Println("direction:", direction)
		}
		checkPos := Vector{pos.x + (direction.x * attempt), pos.y + (direction.y * attempt)}
		if debug {
			fmt.Printf("pos: %+v ", checkPos)
			fmt.Print(string(seatingArea[checkPos]), " ")
		}
		for seatingArea[checkPos] == empty {

			attempt++
			checkPos = Vector{pos.x + (direction.x * attempt), pos.y + (direction.y * attempt)}
			if debug {
				fmt.Printf("pos: %+v ", checkPos)
				fmt.Print(string(seatingArea[checkPos]))
			}
		}

		if seatingArea[checkPos] == takenSeat {
			takenSeats++
		} else if seatingArea[checkPos] == seat {
			seats++
		}
		if debug {
			fmt.Println("")
		}
	}
	if debug {
		fmt.Println("takenSeats:", takenSeats, "seats:", seats)
	}
	if takenSeats >= 5 {
		return seat
	} else if seatingArea[pos] == seat && takenSeats >= 1 {
		return seat
	}
	return takenSeat
}

func printSeatingArea(seatingArea map[Vector]rune, size Vector) {
	for rowPos := 0; rowPos < size.x; rowPos++ {
		for colPos := 0; colPos < size.y; colPos++ {
			pos := Vector{colPos, rowPos}
			fmt.Print(string(seatingArea[pos]))
		}
		fmt.Println("")
	}
	fmt.Println("")
}
