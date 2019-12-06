package main

import (
	"Poker/poker"
	"fmt"
)

func main() {
	game := poker.NewGame(10)

	for i := 0; i < len(game.Hands); i++ {
		fmt.Println(game.Hands[i].Rank)
		fmt.Println(game.Hands[i].Combination.Rank)
		fmt.Println(game.Hands[i].Combination.RelatedCards)
		fmt.Println(game.Hands[i].Combination.UnrelatedCards)
		fmt.Println("")
	}
}
