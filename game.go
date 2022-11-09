package main

const (
	FOUNDATION int = iota
	TABLEAU
	STOCK
)

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

func copyAppend[T any](slice []T, elems ...T) []T {
	size := len(slice) + len(elems)
	out := make([]T, 0, size)
	copy(out, slice)
	return append(out, elems...)
}
