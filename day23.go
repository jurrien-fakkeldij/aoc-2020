package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Day23()
}

type intSlice []int

type CupGame struct {
	cupsMap map[int]*ring.Ring
	cups    *ring.Ring
	maxCup  int
	minCup  int
}

func Day23() {
	fmt.Println("==============  Day23 =============")
	fmt.Println("==============  TEST  =============")
	testGame := setupCupGame("389125467", len("389125467"))

	output, _ := testGame.playGame(10)
	if output == "92658374" {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	testGame2 := setupCupGame("389125467", len("389125467"))
	output, _ = testGame2.playGame(100)
	if output == "67384529" {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT =============")
	labels, _ := setupCupGame("156794823", len("156794823")).playGame(100)
	fmt.Println("Labels:", labels)
	labels, guess := setupCupGame("156794823", 1_000_000).playGame(10_000_000)
	fmt.Println("Guess:", guess)
}

func (game *CupGame) playGame(moves int) (string, int64) {
	numberOfCups := game.cups.Len()
	fmt.Println("cups:", numberOfCups, game.cups)

	for move := 1; move <= moves; move++ {
		removed := game.cups.Unlink(3)
		destinationCup := (numberOfCups+game.cups.Value.(int)-2)%numberOfCups + 1

		removedMap := map[int]bool{}
		for i := 0; i < 3; i++ {
			removedMap[removed.Value.(int)] = true
			removed = removed.Next()
		}

		for removedMap[destinationCup] {
			destinationCup = (numberOfCups+destinationCup-2)%numberOfCups + 1
		}
		game.cupsMap[destinationCup].Link(removed)
		game.cups = game.cups.Next()
	}

	if numberOfCups > 100 {
		return "", game.multiplyAround(1)
	}

	return game.findOrderAfterNumber(1), -1
}

func (game *CupGame) multiplyAround(number int) int64 {
	return int64(game.cupsMap[number].Next().Value.(int)) * int64(game.cupsMap[number].Move(2).Value.(int))
}

func (game *CupGame) findOrderAfterNumber(number int) string {
	out := strings.Builder{}
	numberOfCups := game.cups.Len()
	game.cupsMap[number].Unlink(numberOfCups - 1).Do(func(cup interface{}) {
		out.WriteString(strconv.Itoa(cup.(int)))
	})
	fmt.Println("")

	return out.String()
}

func setupCupGame(gameData string, numberOfCups int) *CupGame {
	game := &CupGame{}

	cups := ring.New(numberOfCups)
	cupsMap := map[int]*ring.Ring{}
	for index := 1; index <= numberOfCups; index++ {
		if cups.Value = index; index <= len(gameData) {
			cups.Value = int(gameData[index-1] - '0')
		}
		cupsMap[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	game.cups = cups
	game.cupsMap = cupsMap

	return game
}
