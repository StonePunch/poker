package poker

import (
	"sort"
)

// Hand definition
type Hand struct {
	Cards       []Card // Needs to be 5 cards
	Combination Combination
}

func sortCardsByValue(cards []Card) []Card {
	sort.Slice(cards, func(i int, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	return cards
}

func NewHand(cards []Card) (hand Hand) {
	hand = Hand{
		Cards:       sortCardsByValue(cards),
		Combination: Combination{},
	}

	combination := hand.hasStraightFlush()
	if combination.Combination == StraightFlush {
		hand.Combination = combination
		return hand
	}

	sameValueCombination := hand.hasSameValues()
	if sameValueCombination.Combination == FourOfAKind ||
		sameValueCombination.Combination == FullHouse {
		hand.Combination = sameValueCombination
		return hand
	}

	combination = hand.hasFlush()
	if combination.Combination == Flush {
		hand.Combination = combination
		return hand
	}

	combination = hand.hasStraight()
	if combination.Combination == Straight {
		hand.Combination = combination
		return hand
	}

	if sameValueCombination.Combination == ThreeOfAKind ||
		sameValueCombination.Combination == TwoPairs ||
		sameValueCombination.Combination == OnePair{
		hand.Combination = sameValueCombination
		return hand
	}

	highCard := hand.Cards[len(hand.Cards) - 1]
	hand.Combination = Combination{
		Combination:    HighCard,
		RelatedCards:   [][]Card{{highCard}},
		UnrelatedCards: hand.Cards[:len(hand.Cards)-1],
		HighestCard:    highCard,
	}

	return hand
}

func (hand Hand) hasStraightFlush() (combination Combination) {
	combination = hand.hasFlush()
	if combination.Combination != Flush {
		return combination
	}

	combination = hand.hasStraight()
	if combination.Combination != Straight {
		return combination
	}

	combination.Combination = StraightFlush

	return combination
}

func (hand Hand) hasFlush() (combination Combination) {
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i - 1].Suit != hand.Cards[i].Suit {
			return combination
		}
	}

	combination = Combination{
		Combination:    Flush,
		RelatedCards:   [][]Card{hand.Cards},
		UnrelatedCards: []Card{},
		HighestCard:    hand.Cards[len(hand.Cards) - 1],
	}

	return combination
}

func (hand Hand) hasStraight() (combination Combination) {
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i - 1].Value + 1 != hand.Cards[i].Value {
			return combination
		}
	}

	combination = Combination{
		Combination:    Straight,
		RelatedCards:   [][]Card{hand.Cards},
		UnrelatedCards: []Card{},
		HighestCard:    hand.Cards[len(hand.Cards) - 1],
	}

	return combination
}

func (hand Hand) hasSameValues() (combination Combination) {
	var pairs [][]Card
	var unrelatedCards []Card
	pairCard := hand.Cards[0]
	pairCounter := 0
	for i := 1; i < len(hand.Cards); i++ {
		if pairCard.Value == hand.Cards[i].Value {
			pairCounter++
		} else if pairCounter != 0 {
			pairs = append(pairs, []Card{})
			for c := 0; c < pairCounter + 1; c++ {
				pairs[len(pairs) - 1] = append(pairs[len(pairs) - 1], pairCard)
			}

			pairCounter = 0
			pairCard = hand.Cards[i]
		} else {
			unrelatedCards = append(unrelatedCards, pairCard)
			pairCard = hand.Cards[i]
		}
	}

	if pairCounter != 0 {
		pairs = append(pairs, []Card{})
		for c := 0; c < pairCounter + 1; c++ {
			pairs[len(pairs) - 1] = append(pairs[len(pairs) - 1], pairCard)
		}
	} else {
		unrelatedCards = append(unrelatedCards, pairCard)
	}

	if pairs == nil {
		return combination
	}

	var rank CombinationRank
	if len(pairs) == 2 {
		rank = FullHouse
		if len(pairs[0]) > len(pairs[1]) {
			rank = FullHouse
			pairs = [][]Card{
				pairs[1],
				pairs[0],
			}
		} else if len(pairs[0]) == len(pairs[1]) {
			rank = TwoPairs
		}
	} else {
		switch len(pairs[0]) {
		case 2:
			rank = OnePair
			break
		case 3:
			rank = ThreeOfAKind
			break
		case 4:
			rank = FourOfAKind
			break
		}
	}

	combination = Combination{
		Combination:    rank,
		RelatedCards:   pairs,
		UnrelatedCards: unrelatedCards,
		HighestCard:    pairs[len(pairs) - 1][0],
	}

	return combination
}

func (hand Hand) hasHighCard() (combination Combination) {

	// TODO: Finish this

	return combination
}
