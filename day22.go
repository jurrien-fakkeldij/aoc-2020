package main

import (
	"aoc-2020/internal/filereader"
	"aoc-2020/internal/transformer"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Day22()
}

type Player struct {
	name          string
	deck          []int
	previousDecks []string
}

type Game struct {
	players []*Player
}

func Day22() {
	fmt.Println("==============  Day22 =============")
	fmt.Println("==============  TEST  =============")
	testGame := setupGame(filereader.ReadFile("./input/day-22/test.txt"))

	output, _ := testGame.playGame()

	if output == 291 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT =============")
	game := setupGame(filereader.ReadFile("./input/day-22/input.txt"))
	score, _ := game.playGame()
	fmt.Println("Score:", score)
}

func (game *Game) playGame() (int, *Player) {
	player1 := game.players[0]
	player2 := game.players[1]

	for len(player1.deck) > 0 && len(player2.deck) > 0 {
		deck1string := strings.Join(transformer.SliceItoa(player1.deck), "|")
		deck2string := strings.Join(transformer.SliceItoa(player2.deck), "|")

		if deckAlreadyPlayed(player1, deck1string) || deckAlreadyPlayed(player2, deck2string) {
			fmt.Println("Deck played already:", player1.deck, player2.deck)
			return determineScore(player1), player1
		}

		player1.previousDecks = append(player1.previousDecks, strings.Join(transformer.SliceItoa(player1.deck), "|"))
		player2.previousDecks = append(player2.previousDecks, strings.Join(transformer.SliceItoa(player2.deck), "|"))

		var p1, p2 int
		p1, player1.deck = player1.deck[0], player1.deck[1:]
		p2, player2.deck = player2.deck[0], player2.deck[1:]

		if len(player1.deck) >= p1 && len(player2.deck) >= p2 {
			fmt.Println("starting new game:", p1, p2, player1.deck, player2.deck, player1.deck[0:p1], player2.deck[0:p2])
			//play subgame
			player1Deck := make([]int, len(player1.deck))
			player2Deck := make([]int, len(player2.deck))
			copy(player1Deck, player1.deck)
			copy(player2Deck, player2.deck)

			newGame := &Game{}
			newGame.players = []*Player{
				&Player{
					name:          player1.name,
					deck:          player1Deck[0:p1],
					previousDecks: []string{}},
				&Player{
					name:          player2.name,
					deck:          player2Deck[0:p2],
					previousDecks: []string{}}}

			_, winningPlayer := newGame.playGame()
			fmt.Println("GameWinner:", winningPlayer.name, player1.deck, player2.deck)
			if winningPlayer.name == player1.name {
				player1.deck = append(player1.deck, p1)
				player1.deck = append(player1.deck, p2)
			} else {
				player2.deck = append(player2.deck, p2)
				player2.deck = append(player2.deck, p1)
			}
		} else {
			if p1 > p2 {
				player1.deck = append(player1.deck, p1)
				player1.deck = append(player1.deck, p2)
			} else {
				player2.deck = append(player2.deck, p2)
				player2.deck = append(player2.deck, p1)
			}
		}
	}

	var winningPlayer *Player
	//determine winner
	if len(player1.deck) != 0 {
		winningPlayer = player1
	} else {
		winningPlayer = player2
	}

	//calculate score
	return determineScore(winningPlayer), winningPlayer
}

func deckAlreadyPlayed(player *Player, deck string) bool {
	for _, deckPlayed := range player.previousDecks {
		if deckPlayed == deck {
			return true
		}
	}
	return false
}

func determineScore(winningPlayer *Player) int {
	score := 0
	calculatingIndex := 1
	for index := len(winningPlayer.deck) - 1; index >= 0; index-- {
		score += winningPlayer.deck[index] * calculatingIndex
		calculatingIndex++
	}
	return score
}

func setupGame(gameData []string) *Game {
	game := &Game{}
	var player *Player
	for _, line := range gameData {
		if line == "" {
			//save player
			game.players = append(game.players, player)
		} else if strings.HasPrefix(line, "Player ") {
			player = &Player{name: line[:len(line)-1]}
			player.deck = []int{}
			player.previousDecks = []string{}
		} else {
			card, _ := strconv.Atoi(line)
			player.deck = append(player.deck, card)
		}
	}

	//save last player as well
	game.players = append(game.players, player)

	return game
}
