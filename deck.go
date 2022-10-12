package impatience

struct Deck {
  Category DeckName
  First *Card
  Last *Card
}

type DeckName int
const (
  Stock = Location iota
  Foundations
  Tableau1
  Tableau2
  Tableau3
  Tableau4
  Tableau5
  Tableau6
  Tableau7
)
