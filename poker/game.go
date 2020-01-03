package poker

import (
	"math/rand"
	"sort"
	"time"
)

type Game struct {
	Players int
	Hands []Hand
}

func NewGame(players int) Game {
	var game Game

	if players > 10 {
		return game
	}

	hands := generateHands()

	for i := 0; i < players; i++ {
		game.Hands = append(game.Hands, hands[i])
	}

	game.Hands = sortHandsByRank(game.Hands)

	return game
}

func generateHands() []Hand  {
	var hands []Hand

	deck := generateDeck()

	handNumber := 0
	cards := make(map[int][]Card)
	for ok := true; ok; ok = len(deck) > 0 {
		random := random(0, len(deck) - 1)

		cards[handNumber] = append(cards[handNumber], deck[random])
		deck = remove(deck, random)

		if len(cards[handNumber]) == 5 {
			handNumber++
		}
	}

	// The last key of the map will only contain 2 cards
	delete(cards, 10)

	for _, card := range cards {
		hands = append(hands, NewHand(card))
	}
	return hands
}

func generateDeck() []Card  {
	var deck []Card

	// Generate a 52 card deck
	const maxSuit Suit = 4
	const maxValue Value = 13
	var suit Suit
	var value Value
	for suit = 1; suit <= maxSuit; suit++ {
		for value = 1; value <= maxValue; value++ {
			deck = append(deck, Card{
				Value: value,
				Suit:  suit,
			})
		}
	}

	return deck
}

func random(min int, max int) int  {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}

func remove(deck []Card, i int) []Card {
	deck[i] = deck[len(deck) - 1]
	return deck[:len(deck) - 1]
}

func sortHandsByRank(hands []Hand) []Hand {
	sort.Slice(hands, func(i int, j int) bool {
		return hands[i].Rank < hands[j].Rank
	})
	return hands
}
