package main

import (
	"Poker/poker"
	"fmt"
)

func main() {
	cards := []poker.Card{
		{
			Value: poker.Three,
			Suit:  poker.Clubs,
		},
		{
			Value: poker.Three,
			Suit:  poker.Spades,
		},
		{
			Value: poker.Three,
			Suit:  poker.Hearts,
		},
		{
			Value: poker.Two,
			Suit:  poker.Spades,
		},
		{
			Value: poker.Four,
			Suit:  poker.Diamonds,
		},
	}

	hand := poker.NewHand(cards)
	fmt.Println(hand.Combination)
}
