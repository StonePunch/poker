package common

import (
	"Poker/src"
	"sort"
)

func SortCardsByValue(cards []poker.Card) []poker.Card {
	sort.Slice(cards, func(i int, j int) bool {
		return cards[i].Face < cards[j].Face
	})

	return cards
}