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

	combination, isNil := hand.hasSameValues()

	combination, isNil = hand.hasStraightFlush()
	if !isNil {
		hand.Combination = combination
		return hand
	}

	// TODO: FourOfAKind

	// TODO: FullHouse

	combination, isNil = hand.hasFlush()
	if !isNil {
		hand.Combination = combination
		return hand
	}

	combination, isNil = hand.hasStraight()
	if !isNil {
		hand.Combination = combination
		return hand
	}

	// TODO: ThreeOfAKind

	// TODO: TwoPairs

	// TODO: OnePair

	// TODO: HighCard

	return hand
}

func (hand Hand) hasStraightFlush() (combination Combination, isNil bool) {
	combination, isNil = hand.hasFlush()
	if isNil {
		return combination, true
	}

	combination, isNil = hand.hasStraight()
	if isNil {
		return combination, true
	}

	combination.Combination = StraightFlush

	return combination, false
}

func (hand Hand) hasFlush() (combination Combination, isNil bool) {
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i-1].Suit != hand.Cards[i].Suit {
			return combination, true
		}
	}

	combination = Combination{
		Combination:    Flush,
		RelatedCards:   [][]Card{hand.Cards},
		UnrelatedCards: []Card{},
		HighestCard:    hand.Cards[len(hand.Cards)-1],
	}

	return combination, false
}

func (hand Hand) hasStraight() (combination Combination, isNil bool) {
	for i := 1; i < len(hand.Cards)-1; i++ {
		if hand.Cards[i-1].Value+1 != hand.Cards[i].Value {
			return combination, true
		}
	}

	combination = Combination{
		Combination:    Straight,
		RelatedCards:   [][]Card{hand.Cards},
		UnrelatedCards: []Card{},
		HighestCard:    hand.Cards[len(hand.Cards)-1],
	}

	return combination, false
}

func (hand Hand) hasSameValues() (combination Combination, isNil bool) {
	var pairs [][]Card
	var unrelatedCards []Card
	pairCard := hand.Cards[0]
	pairCounter := 0
	// TODO: Check wtf is wrong with this 'for' validation
	for i := 1; i <= len(hand.Cards)-1; i++ {
		if pairCard.Value == hand.Cards[i].Value {
			pairCounter++
		} else if pairCounter != 0 {
			pairsLength := len(pairs)
			pairs = append(pairs, []Card{})
			for c := 0; c < pairCounter+1; c++ {
				pairs[pairsLength] = append(pairs[pairsLength], pairCard)
			}

			pairCounter = 0
			pairCard = hand.Cards[i]
		} else {
			unrelatedCards = append(unrelatedCards, pairCard)
			pairCard = hand.Cards[i]
		}
	}

	if pairCounter != 0 {
		pairsLength := len(pairs)
		pairs = append(pairs, []Card{})
		for c := 0; c < pairCounter+1; c++ {
			pairs[pairsLength] = append(pairs[pairsLength], pairCard)
		}
	} else {
		unrelatedCards = append(unrelatedCards, pairCard)
	}

	if pairs == nil {
		return combination, true
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
		HighestCard:    pairs[len(pairs)-1][0],
	}

	return combination, false
}

func (hand Hand) hasHighCard() (combination Combination, isNil bool) {

	// TODO: Finish this

	return combination, false
}
