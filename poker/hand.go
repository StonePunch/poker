package poker

import (
	"fmt"
	"sort"
	"strconv"
)

// Hand definition
type Hand struct {
	Cards       []Card // Needs to be 5 cards
	Combination Combination
	Rank 		int
}

func sortCardsByValue(cards []Card) []Card {
	sort.Slice(cards, func(i int, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	return cards
}

func NewHand(cards []Card) Hand {
	var hand Hand

	if len(cards) != 5 {
		return hand
	}

	hand = Hand{
		Cards:       sortCardsByValue(cards),
		Combination: Combination{},
		Rank: 		 0,
	}

	combination := hand.hasStraightFlush()
	if combination.Rank == StraightFlush {
		hand.Combination = combination
	}

	sameValueCombination := hand.hasSameValues()
	if sameValueCombination.Rank == FourOfAKind ||
		sameValueCombination.Rank == FullHouse {
		hand.Combination = sameValueCombination
	}

	combination = hand.hasFlush()
	if combination.Rank == Flush {
		hand.Combination = combination
	}

	combination = hand.hasStraight()
	if combination.Rank == Straight {
		hand.Combination = combination
	}

	if sameValueCombination.Rank == ThreeOfAKind ||
		sameValueCombination.Rank == TwoPairs ||
		sameValueCombination.Rank == OnePair {
		hand.Combination = sameValueCombination
	}

	if hand.Combination.Rank == 0 {
		highCard := hand.Cards[len(hand.Cards) - 1]
		hand.Combination = Combination{
			Rank:                HighCard,
			RelatedCards:        [][]Card{{highCard}},
			UnrelatedCards:      hand.Cards[:len(hand.Cards) - 1],
		}
	}

	hand.Rank = hand.calcHandRank()

	return hand
}

func (hand Hand) hasStraightFlush() Combination {
	var combination Combination

	combination = hand.hasFlush()
	if combination.Rank != Flush {
		return combination
	}

	combination = hand.hasStraight()
	if combination.Rank != Straight {
		return combination
	}

	combination.Rank = StraightFlush

	return combination
}

func (hand Hand) hasFlush() Combination {
	var combination Combination

	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i - 1].Suit != hand.Cards[i].Suit {
			return combination
		}
	}

	combination = Combination{
		Rank:                Flush,
		RelatedCards:        [][]Card{hand.Cards},
		UnrelatedCards:      []Card{},
	}

	return combination
}

func (hand Hand) hasStraight() Combination {
	var combination Combination

	if hand.Cards[len(hand.Cards) - 1].Value == Ace {
		if hand.Cards[0].Value != Two {
			return combination
		}
		
		for i := 2; i < len(hand.Cards) - 1; i++ {
			if hand.Cards[i - 1].Value + 1 != hand.Cards[i].Value {
				return combination
			}
		}
	} else {
		for i := 1; i < len(hand.Cards); i++ {
			if hand.Cards[i - 1].Value + 1 != hand.Cards[i].Value {
				return combination
			}
		}
	}
	
	combination = Combination{
		Rank:                Straight,
		RelatedCards:        [][]Card{hand.Cards},
		UnrelatedCards:      []Card{},
	}

	return combination
}

// hasSameValues handles the FourOfAKind, FullHouse, ThreeOfAKind, TwoPair and OnePair combinations
func (hand Hand) hasSameValues() Combination {
	var combination Combination

	cardStacks := make(map[Value][]Card)
	for _, card := range hand.Cards {
		cardStacks[card.Value] = append(cardStacks[card.Value], card)
	}

	var pairs [][]Card
	var unrelatedCards []Card
	for _, card := range cardStacks {
		if len(card) > 1 {
			pairs = append(pairs, card)
		} else {
			unrelatedCards = append(unrelatedCards, card[0])
		}
	}

	if pairs == nil {
		return combination
	}

	var rank CombinationRank
	if len(pairs) == 2 {
		rank = FullHouse
		if len(pairs[0]) > len(pairs[1]) {
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
		Rank:                rank,
		RelatedCards:        pairs,
		UnrelatedCards:      unrelatedCards,
	}

	return combination
}

func (hand Hand) calcHandRank() int {
	var rank int

	rankString := fmt.Sprintf("%02d", hand.Combination.Rank)

	lastRelatedSliceIndex := len(hand.Combination.RelatedCards) - 1
	lastRelatedSlice := hand.Combination.RelatedCards[lastRelatedSliceIndex]
	for i := len(lastRelatedSlice) - 1; 0 <= i ; i-- {
		rankString += fmt.Sprintf("%02d", lastRelatedSlice[i].Value)
	}

	relatedCards := hand.Combination.RelatedCards
	if len(relatedCards) == 2 {
		for i := len(relatedCards[0]) - 1; 0 <= i ; i-- {
			rankString += fmt.Sprintf("%02d", relatedCards[0][i].Value)
		}
	}

	unrelatedCards := hand.Combination.UnrelatedCards
	for i := len(unrelatedCards) - 1; 0 <= i; i-- {
		rankString += fmt.Sprintf("%02d", unrelatedCards[i].Value)
	}

	rank, _ = strconv.Atoi(rankString)

	return rank
}