package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

type Ship struct {
	pos      Vector
	waypoint Vector
	memory   map[int]string
}

var NORTH = Vector{0, 1}
var SOUTH = Vector{0, -1}
var WEST = Vector{-1, 0}
var EAST = Vector{1, 0}

type vectorSlice []Vector

func (slice vectorSlice) pos(value Vector) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func (ship *Ship) initializeDockProgram(initProgram []string) int64 {
	ship.memory = make(map[int]string)
	mask := ""

	for _, programLine := range initProgram {
		if strings.Split(programLine, " = ")[0] == "mask" {
			mask = strings.Split(programLine, " = ")[1]
		} else {
			/*fmt.Println("programLine:", programLine)
			fmt.Println("0:", strings.Split(programLine, " = ")[0][4:strings.Index(strings.Split(programLine, " = ")[0], "]")])*/
			memoryAddress, _error := strconv.Atoi(strings.Split(programLine, " = ")[0][4:strings.Index(strings.Split(programLine, " = ")[0], "]")])

			if _error != nil {
				fmt.Println("Error occured reading memoryAddress:", _error)
			}

			data, _error := strconv.Atoi(strings.Split(programLine, " = ")[1])

			if _error != nil {
				fmt.Println("Error occured reading data:", _error)
			}
			ship.setMemory(memoryAddress, data, mask)
		}
	}

	output := int64(0)
	for _, memValue := range ship.memory {
		if i, err := strconv.ParseInt(memValue, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			output += i
		}
	}

	return output
}

func (ship *Ship) initializeDockProgramv2(initProgram []string) int64 {
	ship.memory = make(map[int]string)
	mask := ""

	for _, programLine := range initProgram {
		if strings.Split(programLine, " = ")[0] == "mask" {
			mask = strings.Split(programLine, " = ")[1]
		} else {
			/*fmt.Println("programLine:", programLine)
			fmt.Println("0:", strings.Split(programLine, " = ")[0][4:strings.Index(strings.Split(programLine, " = ")[0], "]")])*/
			memoryAddress, _error := strconv.Atoi(strings.Split(programLine, " = ")[0][4:strings.Index(strings.Split(programLine, " = ")[0], "]")])

			if _error != nil {
				fmt.Println("Error occured reading memoryAddress:", _error)
			}

			data, _error := strconv.Atoi(strings.Split(programLine, " = ")[1])

			if _error != nil {
				fmt.Println("Error occured reading data:", _error)
			}
			ship.setMemoryv2(memoryAddress, data, mask)
		}
	}

	output := int64(0)
	fmt.Println("memory:", ship.memory)
	for _, memoryValue := range ship.memory {
		if i, err := strconv.ParseInt(memoryValue, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			output += i
		}
	}

	return output
}

//00000000000000000000000000000
//000000

func (ship *Ship) setMemoryv2(address int, data int, mask string) {
	bitDataRepresentation := Reverse(strconv.FormatInt(int64(address), 2))
	fmt.Println("data:", bitDataRepresentation)
	fmt.Println("data:", Reverse(mask))

	dataToInsert := []rune{}

	for index, _rune := range Reverse(mask) {
		if _rune == '0' {
			//fmt.Println("index:", index, "length:", len(bitDataRepresentation))
			if index > len(bitDataRepresentation)-1 {
				dataToInsert = append(dataToInsert, '0')
			} else {
				dataToInsert = append(dataToInsert, rune(bitDataRepresentation[index]))
			}
		} else {

			dataToInsert = append(dataToInsert, _rune)
		}
	}
	fmt.Println("data:", string(dataToInsert))
	bitData := string(Reverse(string(dataToInsert)))
	allPossibleMemoryAddresses := computeAllPossibleValues(bitData)
	for _, memaddrValue := range allPossibleMemoryAddresses {

		if i, err := strconv.ParseInt(memaddrValue, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			ship.memory[int(i)] = strconv.FormatInt(int64(data), 2)
		}
	}
	//fmt.Println("%+v", ship.memory)
}

func (ship *Ship) setMemory(address int, data int, mask string) {
	bitDataRepresentation := Reverse(strconv.FormatInt(int64(data), 2))
	//fmt.Println("data:", bitDataRepresentation)
	//fmt.Println("data:", Reverse(mask))

	dataToInsert := []rune{}

	for index, _rune := range Reverse(mask) {
		if _rune == 'X' {
			//fmt.Println("index:", index, "length:", len(bitDataRepresentation))
			if index > len(bitDataRepresentation)-1 {
				dataToInsert = append(dataToInsert, '0')
			} else {
				dataToInsert = append(dataToInsert, rune(bitDataRepresentation[index]))
			}
		} else {
			dataToInsert = append(dataToInsert, _rune)
		}
	}
	//fmt.Println("data ins:", string(Reverse(string(dataToInsert))))
	ship.memory[address] = string(Reverse(string(dataToInsert)))
}

func computeAllPossibleValues(data string) []string {
	output := []string{}
	localOutput := []string{}
	localOutput = append(localOutput, strings.Replace(data, "X", "0", 1))
	localOutput = append(localOutput, strings.Replace(data, "X", "1", 1))

	for _, outputString := range localOutput {
		if strings.Contains(outputString, "X") {
			output = append(output, computeAllPossibleValues(outputString)...)
		} else {
			output = append(output, outputString)
		}
	}

	return output
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (ship *Ship) translateInstruction(instruction string) {
	action := []rune(instruction)[0]
	amount, _error := strconv.Atoi(instruction[1:])
	if _error != nil {
		fmt.Println("Error occured reading instruction:", _error)
		return
	}
	switch action {
	case 'N':
		ship.moveWaypoint(NORTH, amount)
	case 'S':
		ship.moveWaypoint(SOUTH, amount)
	case 'E':
		ship.moveWaypoint(EAST, amount)
	case 'W':
		ship.moveWaypoint(WEST, amount)
	case 'L':
		ship.rotateWaypoint(360 - amount)
	case 'R':
		ship.rotateWaypoint(0 + amount)
	case 'F':
		ship.move(ship.waypoint, amount)
	}
}

func (ship *Ship) translateGuessedInstructions(instruction string) {
	action := []rune(instruction)[0]
	amount, _error := strconv.Atoi(instruction[1:])
	if _error != nil {
		fmt.Println("Error occured reading instruction:", _error)
		return
	}
	switch action {
	case 'N':
		ship.move(NORTH, amount)
	case 'S':
		ship.move(SOUTH, amount)
	case 'E':
		ship.move(EAST, amount)
	case 'W':
		ship.move(WEST, amount)
	case 'L':
		ship.rotateWaypoint(360 - amount)
	case 'R':
		ship.rotateWaypoint(0 + amount)
	case 'F':
		ship.move(ship.waypoint, amount)
	}
}

func (ship *Ship) rotateWaypoint(degrees int) {
	fmt.Println("old waypoint:", ship.waypoint, "degrees:", degrees)

	angle := float64(-degrees) * math.Pi / 180
	fmt.Println("radian:", angle)
	x2 := (float64(ship.waypoint.x) * math.Cos(angle)) - (float64(ship.waypoint.y) * math.Sin(angle))
	y2 := (float64(ship.waypoint.x) * math.Sin(angle)) + (float64(ship.waypoint.y) * math.Cos(angle))
	fmt.Println("x2, y2:", math.Round(x2), math.Round(y2))

	ship.waypoint.x = int(math.Round(x2))
	ship.waypoint.y = int(math.Round(y2))

	fmt.Println("new waypoint", ship.waypoint)
}

func (ship *Ship) moveWaypoint(direction Vector, length int) {
	ship.waypoint.x = ship.waypoint.x + (direction.x * length)
	ship.waypoint.y = ship.waypoint.y + (direction.y * length)
	fmt.Println("new waypoint pos:", ship.waypoint.x, ship.waypoint.y)
}

func (ship *Ship) move(direction Vector, length int) {
	ship.pos.x = ship.pos.x + (direction.x * length)
	ship.pos.y = ship.pos.y + (direction.y * length)
	fmt.Println("new pos:", ship.pos.x, ship.pos.y)
}
