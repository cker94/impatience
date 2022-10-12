package impatience

struct StockDeck {
  List *Deck
  Pos uint8
  LoopCount uint8
}


// To save memory, only deep copy lists that will change.
struct Game {
  Stock StockDeck
  Tableau [7]*Deck
  Foundations [4]*Deck
  PrevMoves []*Moves
  NextMoves []*Moves
}

func NewGame() *Game {
  var game Game
  game.Stock.List = new(Deck)
  var game Game
  game.Foundations = [4]*Deck{
    []Card{}, []Card{}, []Card{}, []Card{}
  }
  game.Tableau = [7][]Card{
    []Card{}, []Card{}, []Card{}, []Card{},
    []Card{}, []Card{}, []Card{}
  }
} 

struct Move {
  Subject *Card
  From *Deck
  To *Deck
}

struct State {
  Moves
}
