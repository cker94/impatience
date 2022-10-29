package impatience

// To save memory, only deep copy lists that will change.
type Game struct {
	Stacks    [9][]*Card
	StockPos  uint8
	StockLoop uint8
	PrevMoves []*Move
	NextMoves []*Move
}

func NewGame() *Game {
	var game Game
	for i := 0; i < len(game.Stacks); i++ {
		game.Stacks[i] = make([]*Card, i+1)
	}
	game.PrevMoves = []*Move{}
	game.NextMoves = make([]*Move, 10)
	return &game
}

type Move struct {
	Subject *Card
	From    StackID
	To      StackID
}
