package impatience

type CardRank int

const (
	UNKNOWN_RANK CardRank = iota
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

type CardSuit int

const (
	UNKNOWN_SUIT CardSuit = iota
	SPADES
	CLUBS
	HEARTS
	DIAMONDS
)

type CardColor int

const (
	UNKNOWN_COLOR CardColor = iota
	BLACK
	RED
)

type Card struct {
	Faceup bool
	Rank   CardRank
	Suit   CardSuit
	Color  CardColor
}
