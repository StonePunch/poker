package poker

// Card definition
type Card struct {
	Value Value
	Suit  Suit
}

// Value enum
type Value int

// All the possible values a card can have
const (
	Two = iota + 1
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

func (value Value) String() string {
	switch value {
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return "Invalid Value"
	}
}

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