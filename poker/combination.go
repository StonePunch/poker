package poker

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

// Combination definition
// RelatedCards -> All the cards being used for the combination
// UnrelatedCards -> Cards left unused in the combination
// HighestCard -> The card with the highest value within the RelatedCards
type Combination struct {
	Combination    CombinationRank
	RelatedCards   [][]Card
	UnrelatedCards []Card
	HighestCard    Card
}