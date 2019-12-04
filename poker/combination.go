package poker

// CombinationRank enum
type CombinationRank int

// All the possible combinations
const (
	HighCard = iota + 1
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func (rank CombinationRank) String() string {
	switch rank {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPairs:
		return "TwoPairs"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case StraightFlush:
		return "StraightFlush"
	default:
		return "Invalid Rank"
	}
}

// Rank definition
// RelatedCards -> All the cards being used for the combination
// UnrelatedCards -> Cards left unused in the combination
// HighestCard -> The card with the highest value within the RelatedCards
type Combination struct {
	Rank           CombinationRank
	RelatedCards   [][]Card
	UnrelatedCards []Card
	HighestCard    Card
}
