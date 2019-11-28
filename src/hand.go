package poker

import (
	"Poker/src/common"
	"Poker/src/enum"
)

// Combination definition
// RelatedCards -> All the cards being used for the combination
// UnrelatedCards -> Cards left unused in the combination
// HighestCard -> The card with the highest value within the RelatedCards
type Combination struct {
	Combination    enum.CombinationRank
	RelatedCards   []Card
	UnrelatedCards []Card
	HighestCard    Card
}

// Hand definition
type Hand struct {
	Cards       []Card // Needs to be 5 cards
	Combination Combination
}

func HasFlush(hand Hand) (combination Combination, isNull bool) {
	hand.Cards = common.SortCardsByValue(hand.Cards)

	var	suit = hand.Cards[0].Suit
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i].Suit != suit {
			return combination, true
		}
	}

	combination = Combination{
		enum.Flush,
		hand.Cards,
		[]Card{},
		hand.Cards[len(hand.Cards) - 1],
	}

	return combination, false
}

func HasStraight(hand Hand) (combination Combination, isNull bool) {
	hand.Cards = common.SortCardsByValue(hand.Cards)

	var previousCard = hand.Cards[0]
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i].Face != previousCard.Face + 1 {
			return combination, true
		}
	}

	combination = Combination{
		enum.Straight,
		hand.Cards,
		[]Card{},
		hand.Cards[len(hand.Cards) - 1],
	}

	return combination, false
}