package poker

// Combination definition
// RelatedCards -> All the cards being used for the combination
// UnrelatedCards -> Cards left unused in the combination
type Combination struct {
	Rank           CombinationRank
	RelatedCards   [][]Card
	UnrelatedCards []Card
}