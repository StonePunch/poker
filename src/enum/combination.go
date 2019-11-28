package enum

// CombinationRank enum
type CombinationRank int

// All the possible combinations
const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	Poker
	StraightFlush
)
