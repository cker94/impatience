package main

import "testing"

func shouldPanic[I comparable, O comparable](t *testing.T, f func(I) O, a I) {
	defer func() { recover() }()
	f(a)
	t.Error("Should have caused a panic:", a)
}

func mapTestAndPanic[I comparable, O comparable](t *testing.T, f func(I) O, tests map[I]O, bad ...I) {
	for input, expected := range tests {
		if output := f(input); output != expected {
			t.Errorf("Output \"%v\" not equal to expected \"%v\"", output, expected)
		}
	}
	for _, b := range bad {
		shouldPanic[I, O](t, f, b)
	}
}

func TestSuitName(t *testing.T) {
	tests := map[CardSuit]string{
		SPADES:       "spades",
		CLUBS:        "clubs",
		HEARTS:       "hearts",
		DIAMONDS:     "diamonds",
		UNKNOWN_SUIT: "unknown suit",
	}
	bad := []CardSuit{-1, 5}
	mapTestAndPanic[CardSuit, string](t, SuitName, tests, bad...)
}

func TestRankName(t *testing.T) {
	tests := map[CardRank]string{
		ACE:          "ace",
		TWO:          "two",
		THREE:        "three",
		FOUR:         "four",
		FIVE:         "five",
		SIX:          "six",
		SEVEN:        "seven",
		EIGHT:        "eight",
		NINE:         "nine",
		TEN:          "ten",
		JACK:         "jack",
		QUEEN:        "queen",
		KING:         "king",
		UNKNOWN_RANK: "unknown rank",
	}
	bad := []CardRank{-1, 14}
	mapTestAndPanic[CardRank, string](t, RankName, tests, bad...)
}

func TestColorName(t *testing.T) {
	tests := map[CardColor]string{
		BLACK:         "black",
		RED:           "red",
		UNKNOWN_COLOR: "unknown color",
	}
	bad := []CardColor{-1, 3}
	mapTestAndPanic[CardColor, string](t, ColorName, tests, bad...)
}
