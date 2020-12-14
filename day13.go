package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Day13()
}

func Day13() {
	fmt.Println("==============  Day13 =============")
	currentTimeString := filereader.ReadFile("./input/day-13/test.txt")[0]
	schedule := filereader.ReadFile("./input/day-13/test.txt")[1]

	fmt.Println("==============  TEST  ============")
	currentTime, error := strconv.Atoi(currentTimeString)
	if error != nil {
		panic("WRONG INPUT TIME:" + currentTimeString)
	}
	output := findClosestTime(currentTime, schedule)
	if output == 295 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = findFirstTimeScheduleMatches(schedule)
	if output == 1068781 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT ============")
	currentTimeString = filereader.ReadFile("./input/day-13/input.txt")[0]
	schedule = filereader.ReadFile("./input/day-13/input.txt")[1]
	currentTime, error = strconv.Atoi(currentTimeString)
	if error != nil {
		panic("WRONG INPUT TIME:" + currentTimeString)
	}
	fmt.Println("Output:", findClosestTime(currentTime, schedule))
	fmt.Println("Output2:", findFirstTimeScheduleMatches(schedule))
}

func findClosestTime(currentTime int, schedule string) int {
	busLines := strings.Split(schedule, ",")
	earliestBus := currentTime
	lowestLeft := 999999999999999999
	for _, busline := range busLines {
		if busline == "x" {
			continue
		} else {
			busNumber, error := strconv.Atoi(busline)

			if error != nil {
				panic("Wrong INPUT:" + busline)
			}
			fmt.Print(busNumber, ":")
			fmt.Println(currentTime % busNumber)
			if lowestLeft > (busNumber - (currentTime % busNumber)) {
				lowestLeft = (busNumber - (currentTime % busNumber))
				earliestBus = busNumber
			}
		}
	}

	return earliestBus * lowestLeft
}

func findFirstTimeScheduleMatches(schedule string) int {
	schedule = strings.ReplaceAll(schedule, "x", "1")
	var buses []int
	for _, busline := range strings.Split(schedule, ",") {
		busNumber, error := strconv.Atoi(busline)

		if error != nil {
			panic("Wrong INPUT:" + busline)
		}
		buses = append(buses, busNumber)
	}

	timestamp := 1

	for {
		timeToSkipIfNoMatch := 1
		valid := true

		for offset := 0; offset < len(buses); offset++ {

			if (timestamp+offset)%buses[offset] != 0 {

				valid = false
				break
			}
			timeToSkipIfNoMatch *= buses[offset]
		}

		// Did we find a full match?
		if valid {
			return timestamp
		}

		timestamp += timeToSkipIfNoMatch
	}
}
