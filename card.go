package impatience

struct Card {
  Rank CardRank
  Suit CardSuit
  Parent *Deck
  Above *Card
  Below *Card
}

type CardRank int
const (
  Unknown = CardRank iota
  Ace
  Two
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
)

type CardSuit int
const (
  Spades = CardSuit iota
  Clubs
  Hearts
  Diamonds
)