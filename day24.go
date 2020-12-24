package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
)

func main() {
	Day24()
}

type DoubleVector struct {
	x int
	y int
}
type FloorTile struct {
	pos   DoubleVector
	white bool
}

func Day24() {
	fmt.Println("==============  Day24 =============")
	fmt.Println("==============  TEST  =============")
	testTraversal := filereader.ReadFile("./input/day-24/test.txt")

	output := countBlackTilesAfterPaths(testTraversal, 0)
	if output == 10 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = countBlackTilesAfterPaths(testTraversal, 1)
	if output == 15 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = countBlackTilesAfterPaths(testTraversal, 100)
	if output == 2208 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT =============")
	fmt.Println("BlackTiles:", countBlackTilesAfterPaths(filereader.ReadFile("./input/day-24/input.txt"), 0))
	fmt.Println("BlackTilesAfter100days:", countBlackTilesAfterPaths(filereader.ReadFile("./input/day-24/input.txt"), 100))
}

func countBlackTilesAfterPaths(paths []string, afterDays int) int {
	//true  = white
	//false = black
	tiles := map[DoubleVector]FloorTile{}
	//referenceTile := &FloorTile{}
	referencePos := DoubleVector{x: 0, y: 0}
	//tiles[referencePos] = referenceTile
	//referenceTile.createAndLinkNeighbours(tiles)
	//fmt.Println("")
	//fmt.Println(tiles)

	for _, path := range paths {
		currentPos := referencePos
		//fmt.Println("referenceTile:", referenceTile)
		//fmt.Println("tiles:", tiles)
		//fmt.Println(path)
		for curPos := 0; curPos < len(path); curPos++ {
			//fmt.Print(string(path[curPos]))
			x, y := 0, 0
			switch path[curPos] {
			case 'n':
				curPos++
				nextChar := path[curPos]
				//fmt.Print(string(path[curPos]))
				if nextChar == 'e' {
					x = 5
					y = 5
				} else if nextChar == 'w' {
					x = -5
					y = 5
				}
				break
			case 's':
				curPos++
				nextChar := path[curPos]
				//fmt.Print(string(path[curPos]))
				if nextChar == 'e' {
					x = 5
					y = -5
				} else if nextChar == 'w' {
					x = -5
					y = -5
				}
				break
			case 'e':
				x = 10
				break
			case 'w':
				x = -10
				break
			}
			//	fmt.Print(" ")
			//currentTile.createAndLinkNeighbours(tiles)
			currentPos = DoubleVector{x: currentPos.x + x, y: currentPos.y + y}
			//fmt.Println("")
		}
		//fmt.Println("")
		/*fmt.Println("reference tile:", referenceTile)
		fmt.Println("flipping tile:", currentTile)*/
		/*if referenceTile == currentTile {
			fmt.Println("yay")
		}*/
		//fmt.Println("flipping to:", !tiles[currentTile])

		if _, exists := tiles[currentPos]; exists {
			//fmt.Println("exists: ", currentPos)
			tile := tiles[currentPos]
			tile.white = !tile.white
			tiles[currentPos] = tile
		} else {
			//fmt.Println("not exists: ", currentPos)
			tiles[currentPos] = FloorTile{pos: currentPos, white: false}
		}
	}

	blackCount := 0
	//printFloor(tiles)
	if afterDays > 0 {
		for day := 1; day <= afterDays; day++ {
			tiles = determineLivingArt(tiles)
			//printFloor(tiles)
		}
	}

	for _, tile := range tiles {
		if !tile.white {
			blackCount++
		}
	}
	return blackCount
}

func determineLivingArt(tiles map[DoubleVector]FloorTile) map[DoubleVector]FloorTile {
	//printFloor(tiles)
	//selectedTiles := tiles
	newFloorTiles := map[DoubleVector]FloorTile{}

	for pos, tile := range tiles {
		noBlackNeighbours, newNeighbourTiles := determineNeighbours(tiles, pos)
		newTile := FloorTile{pos: pos, white: tile.white}
		if tile.white {
			if noBlackNeighbours == 2 {
				newTile.white = false
				newFloorTiles[pos] = newTile
			}
		} else {
			if noBlackNeighbours == 0 || noBlackNeighbours > 2 {
				newTile.white = true
			} else {
				newFloorTiles[pos] = newTile
			}
		}

		for _, neighbourPos := range newNeighbourTiles {
			neighBourTile := FloorTile{pos: neighbourPos, white: true}
			noBlackNeighbours, _ := determineNeighbours(tiles, neighbourPos)

			if noBlackNeighbours == 2 {
				neighBourTile.white = false
				newFloorTiles[neighbourPos] = neighBourTile
			}
		}
	}
	//fmt.Println("newtiles", newFloorTiles)
	//printFloor(newFloorTiles)
	return newFloorTiles
}

func determineNeighbours(tiles map[DoubleVector]FloorTile, pos DoubleVector) (int, []DoubleVector) {
	blackCount := 0
	newNeighBours := []DoubleVector{}

	neighbours := []DoubleVector{
		DoubleVector{x: pos.x + 5, y: pos.y + 5},
		DoubleVector{x: pos.x + 10, y: pos.y},
		DoubleVector{x: pos.x + 5, y: pos.y + -5},
		DoubleVector{x: pos.x + -5, y: pos.y + -5},
		DoubleVector{x: pos.x + -10, y: pos.y},
		DoubleVector{x: pos.x + -5, y: pos.y + 5}}

	for _, neighbourPos := range neighbours {
		if _, exists := tiles[neighbourPos]; !exists {
			newNeighBours = append(newNeighBours, neighbourPos)
		} else {
			if !tiles[neighbourPos].white {
				blackCount++
			}
		}
	}
	//fmt.Println("tile:", pos, "newNeighbours:", newNeighBours)
	//fmt.Println("")
	return blackCount, newNeighBours
}

func printFloor(tiles map[DoubleVector]FloorTile) {
	for y := -10; y <= 10; y++ {
		for x := -10; x <= 10; x++ {
			pos := DoubleVector{x: 5 * x, y: 5 * y}
			if tile, exists := tiles[pos]; exists {
				if tile.white {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("")
	}
}

/*func (tile *FloorTile) createAndLinkNeighbours(tiles map[*FloorTile]bool) {
	if tile.ne == nil {
		fmt.Println("new ne")
		tile.ne = &FloorTile{}
		defer func() {
			tile.ne.se = tile.e
			tile.ne.sw = tile
			tile.ne.w = tile.nw
			tiles[tile.ne] = true
		}()
	}

	if tile.e == nil {
		tile.e = &FloorTile{}
		fmt.Println("new e")
		defer func() {
			tile.e.nw = tile.ne
			tile.e.w = tile
			tile.e.sw = tile.se
			tiles[tile.e] = true
		}()
	}

	if tile.se == nil {
		fmt.Println("new se")
		tile.se = &FloorTile{}
		defer func() {
			tile.se.nw = tile
			tile.se.w = tile.sw
			tile.se.ne = tile.e
			tiles[tile.se] = true
		}()
	}

	if tile.sw == nil {
		fmt.Println("new sw")
		tile.sw = &FloorTile{}
		defer func() {
			tile.sw.ne = tile
			tile.sw.e = tile.se
			tile.sw.nw = tile.w
			tiles[tile.sw] = true
		}()
	}

	if tile.w == nil {
		fmt.Println("new w")
		tile.w = &FloorTile{}

		defer func() {
			tile.w.ne = tile.nw
			tile.w.e = tile
			tile.w.se = tile.sw
			tiles[tile.w] = true
		}()
	}

	if tile.nw == nil {
		fmt.Println("new nw")
		tile.nw = &FloorTile{}
		defer func() {
			tile.nw.e = tile.ne
			tile.nw.se = tile
			tile.nw.sw = tile.w
			tiles[tile.nw] = true
		}()
	}

	//return tiles
}*/
