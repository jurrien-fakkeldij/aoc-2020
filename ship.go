package main

import (
	"fmt"
	"math"
	"strconv"
)

type Vector struct {
	x int
	y int
}

type Ship struct {
	pos      Vector
	waypoint Vector
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
