package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
)

func main() {
	Day17()
}

type Vector3d struct {
	x int
	y int
	z int
	w int
}

type Cube struct {
	active   bool
	location Vector3d
}

func (cube Cube) getNeighbours(currentSpace map[Vector3d]Cube) []Vector3d {
	neighbourCubes := []Vector3d{}
	for w := -1; w < 2; w++ {
		for z := -1; z < 2; z++ {
			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					location := Vector3d{cube.location.x + x, cube.location.y + y, cube.location.z + z, cube.location.w + w}
					//fmt.Println("adding location:", location)
					if location != cube.location {
						neighbourCubes = append(neighbourCubes, location)
					}
				}
			}
		}
	}

	return neighbourCubes
}

func Day17() {
	fmt.Println("==============  Day17 =============")
	testInputMap := filereader.ReadFile("./input/day-17/test.txt")
	fmt.Println("==============  TEST  =============")

	testSpaceMap := buildInitialSpace(testInputMap)

	activeCubes := countActiveCubesAfterCycles(6, testSpaceMap)

	if activeCubes == 848 {
		fmt.Println("Test succeeded:", activeCubes)
	} else {
		fmt.Println("Test failed:", activeCubes)
	}

	fmt.Println("============== OUTPUT =============")

	inputMap := filereader.ReadFile("./input/day-17/input.txt")
	spaceMap := buildInitialSpace(inputMap)
	fmt.Println("Active cubes after boot:", countActiveCubesAfterCycles(6, spaceMap))
}

func countActiveCubesAfterCycles(cycles int, currentSpace map[Vector3d]Cube) int {
	//printSpaceMap(currentSpace)
	for i := 1; i <= cycles; i++ {
		currentSpace = cycle(currentSpace)
		//fmt.Println("cycle:", i)
		//printSpaceMap(currentSpace)
	}
	activeCubes := 0
	for _, cube := range currentSpace {
		if cube.active {
			activeCubes++
		}
	}

	return activeCubes
}

func cycle(currentSpace map[Vector3d]Cube) map[Vector3d]Cube {
	newSpace := make(map[Vector3d]Cube)

	for location, cube := range currentSpace {
		//fmt.Println("getNeighBours:", cube)
		neighbours := cube.getNeighbours(currentSpace)
		activeNeighbours := 0
		for _, neighbourLoc := range neighbours {
			if _, exists := currentSpace[neighbourLoc]; !exists {
				//fmt.Println("doesnt exist:", neighbourLoc)
				newCube := Cube{false, neighbourLoc}
				newNeighbourActive := 0
				newNeighbours := newCube.getNeighbours(currentSpace)
				for _, newNeighbourLoc := range newNeighbours {
					if _, exists := currentSpace[newNeighbourLoc]; exists {
						if currentSpace[newNeighbourLoc].active {
							newNeighbourActive++
						}
					}
				}
				newCubeActive := false
				if newNeighbourActive == 3 {
					newCubeActive = true
				}
				newSpace[neighbourLoc] = Cube{newCubeActive, neighbourLoc}
			} else {
				if currentSpace[neighbourLoc].active {
					activeNeighbours++
				}
			}
		}
		if cube.active {
			if activeNeighbours == 2 || activeNeighbours == 3 {
				cube.active = true
			} else {
				cube.active = false
			}
		} else {
			if activeNeighbours == 3 {
				cube.active = true
			} else {
				cube.active = false
			}
		}
		newSpace[location] = cube
	}

	return newSpace
}

func buildInitialSpace(inputMap []string) map[Vector3d]Cube {
	spaceMap := make(map[Vector3d]Cube)

	for row, line := range inputMap {
		for column, character := range []rune(line) {
			pos := Vector3d{column, row, 0, 0}
			active := character == '#'
			cube := Cube{active, pos}
			spaceMap[pos] = cube
		}
	}

	return spaceMap
}

/*func printSpaceMap(currentSpace map[Vector3d]Cube) {
	for levelPos := -6; levelPos < 6; levelPos++ {
		fmt.Println("level:", levelPos)
		for rowPos := -6; rowPos < 6; rowPos++ {
			for colPos := -6; colPos < 6; colPos++ {
				pos := Vector3d{colPos, rowPos, levelPos}
				char := '+'
				if _, exists := currentSpace[pos]; exists {
					if currentSpace[pos].active {
						char = '#'
					} else {
						char = '.'
					}
					fmt.Print(string(char))
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}*/
