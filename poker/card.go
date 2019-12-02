package poker

// Value enum
type Value int

// All the possible values a card can have
const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Suit enum
type Suit int

// All the possible suits a card can have
const (
	Hearts = iota
	Spades
	Clubs
	Diamonds
)

// Card definition
type Card struct {
	Value Value
	Suit  Suit
}