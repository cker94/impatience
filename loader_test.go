package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestLoadFile(t *testing.T) {
	var exp SaveData
	inputs := []string{"game.toml", "game.json"}
	// exp stock values.
	exp.Stock.Limit = 3
	exp.Stock.Stack = []string{"sA", "sK", "c2", "cK", "sQ", "s4", "d5", "c3", "c10", "dJ", "cA", "cQ", "h9", "hQ", "d4", "h4", "d3", "hA", "s6", "c4", "d2", "cJ", "h7", "d10"}
	// exp tableau values.
	exp.Tableau.Stacks = [][]string{
		{"d7"},
		{"c9", "h10"},
		{"h2", "s8", "h5"},
		{"c6", "h3", "s9", "c8"},
		{"s5", "c7", "dA", "c5", "s2"},
		{"??", "??", "??", "??", "sJ", "s10"},
		{"hK", "h6", "dK", "d8", "s7", "h8", "d9"},
	}
	exp.Tableau.Facedown = []int{0, 1, 2, 3, 4, 5, 6}
	// exp foundation values.
	exp.Foundations = map[string][]string{
		"spades":   {},
		"clubs":    {},
		"hearts":   {},
		"diamonds": {},
	}
	expected, marsherr := json.Marshal(exp)
	if marsherr != nil {
		t.Error("Setup error:", marsherr)
	}

	// Begin testing.
	for i, input := range inputs {
		out, loaderr := LoadFile(input)
		if loaderr != nil {
			t.Errorf("Test %d: %v", i, loaderr)
			continue
		}
		output, marsherr := json.Marshal(out)
		if marsherr != nil {
			t.Errorf("Test %d: %v", i, marsherr)
			continue
		}
		if bytes.Compare(output, expected) != 0 {
			msg := `Test %d: LoadFile(%q) != expected:

output:   %s

expected: %s
`
			t.Errorf(msg, i, input, output, expected)
		}
	}
}

func TestImport(t *testing.T) {
	var exp Game
	inputs := []string{"game.toml", "game.json"}
	// exp stock values.
	exp.Stock.Limit = 3
	exp.Stock.Stack = []*Card{
		{ACE, SPADES, BLACK}, {KING, SPADES, BLACK}, {TWO, CLUBS, BLACK}, {KING, CLUBS, BLACK}, {QUEEN, SPADES, BLACK}, {FOUR, SPADES, BLACK}, {FIVE, DIAMONDS, RED}, {THREE, CLUBS, BLACK}, {TEN, CLUBS, BLACK}, {JACK, DIAMONDS, RED}, {ACE, CLUBS, BLACK}, {QUEEN, CLUBS, BLACK}, {NINE, HEARTS, RED}, {QUEEN, HEARTS, RED}, {FOUR, DIAMONDS, RED}, {FOUR, HEARTS, RED}, {THREE, DIAMONDS, RED}, {ACE, HEARTS, RED}, {SIX, SPADES, BLACK}, {FOUR, CLUBS, BLACK}, {TWO, DIAMONDS, RED}, {JACK, CLUBS, BLACK}, {SEVEN, HEARTS, RED}, {TEN, DIAMONDS, RED},
	}
	// exp tableau values.
	exp.Tableau.Stacks = [7][]*Card{
		{&Card{SEVEN, DIAMONDS, RED}},
		{&Card{NINE, CLUBS, BLACK}, &Card{TEN, HEARTS, RED}},
		{&Card{TWO, HEARTS, RED}, &Card{EIGHT, SPADES, BLACK}, &Card{FIVE, HEARTS, RED}},
		{&Card{SIX, CLUBS, BLACK}, &Card{THREE, HEARTS, RED}, &Card{NINE, SPADES, BLACK}, &Card{EIGHT, CLUBS, BLACK}},
		{&Card{FIVE, SPADES, BLACK}, &Card{SEVEN, CLUBS, BLACK}, &Card{ACE, DIAMONDS, RED}, &Card{FIVE, CLUBS, BLACK}, &Card{TWO, SPADES, BLACK}},
		{&Card{UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, &Card{UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, &Card{UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, &Card{UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, &Card{JACK, SPADES, BLACK}, &Card{TEN, SPADES, BLACK}},
		{&Card{KING, HEARTS, RED}, &Card{SIX, HEARTS, RED}, &Card{KING, DIAMONDS, RED}, &Card{EIGHT, DIAMONDS, RED}, &Card{SEVEN, SPADES, BLACK}, &Card{EIGHT, HEARTS, RED}, &Card{NINE, DIAMONDS, RED}},
	}
	exp.Tableau.Facedown = [7]int{0, 1, 2, 3, 4, 5, 6}
	// exp foundation values.
	exp.Foundations = [4][]*Card{
		{}, {}, {}, {},
	}

	// Convert to JSON for comparison.
	expected, marsherr := json.Marshal(exp)
	if marsherr != nil {
		t.Error("Setup error:", marsherr)
	}

	// Begin testing.
	for i, input := range inputs {
		var subject Game
		// Load file data
		save, loaderr := LoadFile(input)
		if loaderr != nil {
			t.Errorf("Test %d: %v", i, loaderr)
			continue
		}
		// Import data into gane
		if err := subject.Import(save); err != nil {
			t.Errorf("Test %d: %v", i, loaderr)
			continue
		}
		// Convert to JSON
		output, marsherr := json.Marshal(subject)
		if marsherr != nil {
			t.Errorf("Test %d: %v", i, marsherr)
			continue
		}
		// Compare to expected.
		if bytes.Compare(output, expected) != 0 {
			msg := `Test %d: (Game).Import(%q) != expected:

output:   %s

expected: %s
`
			t.Errorf(msg, i, input, output, expected)
		}
	}
}
