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
			Value: poker.Three,
			Suit:  poker.Hearts,
		},
		{
			Value: poker.Five,
			Suit:  poker.Spades,
		},
		{
			Value: poker.Five,
			Suit:  poker.Diamonds,
		},
	}

	poker.NewHand(cards)
}
