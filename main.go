package main

import (
	"fmt"
	"math/rand"
)

type Ghost struct {
	gameWord string
}

func (g *Ghost) Over(p Player) bool {
	corpus := GetCorpus()
	over := corpus.IsAWord(g.gameWord)
	if over {
		fmt.Printf("[%s won]\n", p.GetType())
	}
	return over
}

// returns true for challenge
func (g *Ghost) Play(p Player) bool {
	letter := p.Play(g.gameWord)
	if letter == 0 {
		return true
	}
	g.gameWord = fmt.Sprintf("%s%s", g.gameWord, string(letter))
	fmt.Println("word:", g.gameWord)
	return false
}

func (g *Ghost) PlayChallenge(primary Player, challenged Player) {
	if challenged.PlayChallenge(g.gameWord) {
		fmt.Printf("[%s won]\n", challenged.GetType())
	} else {
		fmt.Printf("[%s won]\n", primary.GetType())
	}
}

func GetPlayers(players []Player, i int) (active Player, idle Player) {
	if i == 0 {
		active = players[1]
		idle = players[0]
	} else {
		active = players[0]
		idle = players[1]
	}
	return
}

func SwapIndex(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}

func main() {
	ghost := Ghost{}

	players := []Player{
		&Computer{},
		&Human{},
	}

	playerIndex := rand.Intn(1)

	for {
		activePlayer, idlePlayer := GetPlayers(players, playerIndex)

		if ghost.Play(activePlayer) {
			ghost.PlayChallenge(activePlayer, idlePlayer)
			break
		}

		if ghost.Over(idlePlayer) {
			break
		}

		playerIndex = SwapIndex(playerIndex)
	}

}
