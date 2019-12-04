package main

import (
	"Poker/poker"
	"fmt"
)

func main() {
	cards := []poker.Card{
		{
			Value: poker.Three,
			Suit:  poker.Hearts,
		},
		{
			Value: poker.Four,
			Suit:  poker.Spades,
		},
		{
			Value: poker.Ace,
			Suit:  poker.Hearts,
		},
		{
			Value: poker.Two,
			Suit:  poker.Clubs,
		},
		{
			Value: poker.Five,
			Suit:  poker.Hearts,
		},
	}

	hand := poker.NewHand(cards)
	fmt.Println(hand.Combination)
	fmt.Println(hand.Rank)
}
