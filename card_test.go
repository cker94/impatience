package main

import (
	"fmt"
	"testing"
)

var testCards []*Card = []*Card{
	{ACE, SPADES, BLACK},
	{TWO, SPADES, BLACK},
	{THREE, SPADES, BLACK},
	{FOUR, SPADES, BLACK},
	{FIVE, SPADES, BLACK},
	{SIX, SPADES, BLACK},
	{SEVEN, SPADES, BLACK},
	{EIGHT, SPADES, BLACK},
	{NINE, SPADES, BLACK},
	{TEN, SPADES, BLACK},
	{JACK, SPADES, BLACK},
	{QUEEN, SPADES, BLACK},
	{KING, SPADES, BLACK},
	{ACE, CLUBS, BLACK},
	{TWO, CLUBS, BLACK},
	{THREE, CLUBS, BLACK},
	{FOUR, CLUBS, BLACK},
	{FIVE, CLUBS, BLACK},
	{SIX, CLUBS, BLACK},
	{SEVEN, CLUBS, BLACK},
	{EIGHT, CLUBS, BLACK},
	{NINE, CLUBS, BLACK},
	{TEN, CLUBS, BLACK},
	{JACK, CLUBS, BLACK},
	{QUEEN, CLUBS, BLACK},
	{KING, CLUBS, BLACK},
	{ACE, HEARTS, RED},
	{TWO, HEARTS, RED},
	{THREE, HEARTS, RED},
	{FOUR, HEARTS, RED},
	{FIVE, HEARTS, RED},
	{SIX, HEARTS, RED},
	{SEVEN, HEARTS, RED},
	{EIGHT, HEARTS, RED},
	{NINE, HEARTS, RED},
	{TEN, HEARTS, RED},
	{JACK, HEARTS, RED},
	{QUEEN, HEARTS, RED},
	{KING, HEARTS, RED},
	{ACE, DIAMONDS, RED},
	{TWO, DIAMONDS, RED},
	{THREE, DIAMONDS, RED},
	{FOUR, DIAMONDS, RED},
	{FIVE, DIAMONDS, RED},
	{SIX, DIAMONDS, RED},
	{SEVEN, DIAMONDS, RED},
	{EIGHT, DIAMONDS, RED},
	{NINE, DIAMONDS, RED},
	{TEN, DIAMONDS, RED},
	{JACK, DIAMONDS, RED},
	{QUEEN, DIAMONDS, RED},
	{KING, DIAMONDS, RED},

	{TEN, SPADES, BLACK},
	{TEN, CLUBS, BLACK},
	{TEN, HEARTS, RED},
	{TEN, DIAMONDS, RED},
}

var testCodes []string = []string{
	"SA", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10", "SJ", "SQ", "SK",
	"CA", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10", "CJ", "CQ", "CK",
	"HA", "H2", "H3", "H4", "H5", "H6", "H7", "H8", "H9", "H10", "HJ", "HQ", "HK",
	"DA", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "D10", "DJ", "DQ", "DK",
	"S1", "C1", "h1", "d1",
}

// Generic function for testing panics.
func shouldPanic(t *testing.T, f func() string) {
	defer func() { recover() }()
	m := f()
	t.Error(m)
}

func shouldPanicAll[I any, O any](t *testing.T, f func(I) O, a []I) {
	size := len(a)
	for i := 0; i < size; i++ {
		func() {
			defer func() { recover() }()
			f(a[i])
			t.Error("Failed to panic:", a[i])
		}()
	}
}

// Generic function to compare output of single-input, single-output functions with expected output via a map.
func mapTest[I comparable, O comparable](t *testing.T, f func(I) O, tests map[I]O) {
	for input, expected := range tests {
		if output := f(input); output != expected {
			t.Errorf("Output \"%v\" not equal to expected \"%v\"", output, expected)
		}
	}
}

func TestCardId(t *testing.T) {
	inputs := testCards[:len(testCards)-4]
	expected := testCodes[:len(testCodes)-4]
	bad := []*Card{
		{-1, 0, 0},
		{14, 0, 0},
		{0, -1, 0},
		{0, 5, 0},
	}

	// Verify inputs and expected are the same length.
	size := len(inputs)
	if len(expected) != size {
		t.Errorf("Setup error: inputs and expected list size mismatch: %d inputs; %d expected.", len(inputs), len(expected))
	}
	// Test all valid inputs.
	for i := 0; i < size; i++ {
		output := inputs[i].Id()
		if output != expected[i] {
			t.Errorf(`Unexpected output: (inputs[%d]) -> %q; expected %q.`,
				i, output, expected[i])
		}
	}
	// Test invalid inputs.
	size = len(bad)
	for i := 0; i < size; i++ {
		shouldPanic(t, func() string {
			bad[i].Id()
			return fmt.Sprintf("Failed to panic: bad[%d]", i)
		})
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
	mapTest(t, SuitName, tests)
	shouldPanicAll(t, SuitName, bad)
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
	mapTest(t, RankName, tests)
	shouldPanicAll(t, RankName, bad)
}

func TestColorName(t *testing.T) {
	tests := map[CardColor]string{
		BLACK:         "black",
		RED:           "red",
		UNKNOWN_COLOR: "unknown color",
	}
	bad := []CardColor{-1, 3}
	mapTest(t, ColorName, tests)
	shouldPanicAll(t, ColorName, bad)
}

func TestParseCard(t *testing.T) {
	inputs := testCodes
	expected := testCards
	bad := []string{"S0", "cL", "h11", "gA", "gab", "321"}

	// Verify inputs and expected are the same length.
	size := len(inputs)
	if len(expected) != size {
		t.Errorf("Setup error: inputs and expected list size mismatch: %d inputs; %d expected.", len(inputs), len(expected))
	}
	// Test all valid inputs.
	for i := 0; i < size; i++ {
		output, err := ParseCard(inputs[i])
		if err != nil {
			t.Errorf("(%q)-> Error: %v", inputs[i], err)
		}
		if *output != *expected[i] {
			t.Errorf(`Unexpected output: (%q) ->
           *Card{Rank: %d, Suit: %d, Color: %d}
expected *Card{Rank: %d, Suit: %d, Color: %d}`,
				inputs[i], output.Rank, output.Suit, output.Color,
				expected[i].Rank, expected[i].Suit, expected[i].Color)
		}
	}
	// Test invalid inputs.
	for _, b := range bad {
		output, err := ParseCard(b)
		if err == nil {
			t.Errorf("Expected error from: (%q)-> output = *Card{Rank: %d, Suit: %d, Color: %d}",
				b, output.Rank, output.Suit, output.Color)
		}
	}
}

func TestParseCards(t *testing.T) {
	inputs := testCodes
	expected := testCards
	bad := []string{"S0", "cL", "h11", "gA", "gab", "321"}

	// Verify inputs and expected are the same length.
	size := len(inputs)
	if len(expected) != size {
		t.Errorf("Setup error: inputs and expected list size mismatch: %d inputs; %d expected.", len(inputs), len(expected))
	}
	// Test valid input.
	output, err := ParseCards(inputs)
	if err != nil {
		t.Errorf("(testCodes []string)-> Error: %v", err)
	}
	for i := 0; i < size; i++ {
		if *output[i] != *expected[i] {
			t.Errorf(`Unexpected output: (%q) ->
             *Card{Rank: %d, Suit: %d, Color: %d}
  expected = *Card{Rank: %d, Suit: %d, Color: %d}`,
				inputs[i], output[i].Rank, output[i].Suit, output[i].Color,
				expected[i].Rank, expected[i].Suit, expected[i].Color)
		}
	}
	// Test invalid inputs.
	for i := 0; i < len(bad); i++ {
		_, err := ParseCards(bad[i:])
		if err == nil {
			t.Errorf("Expected error from: (bad [%d:]string)-> []*Card", i)
		}
	}
}
