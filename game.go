package main

const (
	FOUNDATION int = iota
	TABLEAU
	STOCK
)

/*const (
	FOUNDATION_1 int = iota
	FOUNDATION_2
	FOUNDATION_3
	FOUNDATION_4
	TABLEAU_1
	TABLEAU_2
	TABLEAU_3
	TABLEAU_4
	TABLEAU_5
	TABLEAU_6
	TABLEAU_7
)*/

type Game struct {
	Foundations [4][]*Card
	Stock       struct {
		Limit int
		Loop  int
		Pos   int
		Stack []*Card
	}
	Tableau struct {
		Facedown [7]int
		Stacks   [7][]*Card
	}
	Moves struct {
		Prev []*Move
		Next []*Move
	}
}

type Move struct {
	Card *Card
	To   struct {
		Category int
		Stack    int
	}
	From struct {
		Category int
		Stack    int
		Index    int
	}
}
