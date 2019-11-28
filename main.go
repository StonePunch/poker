package main

import (
	"Poker/src"
	"Poker/src/enum"
)

func main() {
	hand := poker.Hand{
		Cards: []poker.Card{
			{
				Face: enum.Two,
				Suit: enum.Spades,
			},
			{
				Face: enum.Seven,
				Suit: enum.Spades,
			},
			{
				Face: enum.Four,
				Suit: enum.Spades,
			},
			{
				Face: enum.Queen,
				Suit: enum.Spades,
			},
			{
				Face: enum.Eight,
				Suit: enum.Spades,
			},
		},
	}

	var combination poker.Combination
	var isNull = false

	// TODO: StraightFlush

	// TODO: Poker

	// TODO: FullHouse

	combination, isNull = poker.HasFlush(hand)
	if !isNull {
		hand.Combination = combination
	}

	combination, isNull = poker.HasStraight(hand)
	if !isNull {
		hand.Combination = combination
	}

	// TODO: ThreeOfAKind

	// TODO: TwoPairs

	// TODO: OnePair

	// TODO: HighCard
}
