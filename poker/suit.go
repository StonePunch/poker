package poker

// Suit enum
type Suit int

// All the possible suits a card can have
const (
	Hearts = iota + 1
	Spades
	Clubs
	Diamonds
)

func (suit Suit) String() string {
	switch suit {
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	default:
		return "Invalid Suit"
	}
}