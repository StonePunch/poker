package main

import (
	"Poker/poker"
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
			Value: poker.Two,
			Suit:  poker.Hearts,
		},
		{
			Value: poker.Four,
			Suit:  poker.Spades,
		},
		{
			Value: poker.Four,
			Suit:  poker.Diamonds,
		},
	}

	poker.NewHand(cards)
}
