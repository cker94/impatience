package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

// TODO: Add more variations of toml and json files to load.
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
		t.SkipNow()
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

// TODO: add more variations to test.
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
		{{SEVEN, DIAMONDS, RED}},
		{{NINE, CLUBS, BLACK}, {TEN, HEARTS, RED}},
		{{TWO, HEARTS, RED}, {EIGHT, SPADES, BLACK}, {FIVE, HEARTS, RED}},
		{{SIX, CLUBS, BLACK}, {THREE, HEARTS, RED}, {NINE, SPADES, BLACK}, {EIGHT, CLUBS, BLACK}},
		{{FIVE, SPADES, BLACK}, {SEVEN, CLUBS, BLACK}, {ACE, DIAMONDS, RED}, {FIVE, CLUBS, BLACK}, {TWO, SPADES, BLACK}},
		{{UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, {UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, {UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, {UNKNOWN_RANK, UNKNOWN_SUIT, UNKNOWN_COLOR}, {JACK, SPADES, BLACK}, {TEN, SPADES, BLACK}},
		{{KING, HEARTS, RED}, {SIX, HEARTS, RED}, {KING, DIAMONDS, RED}, {EIGHT, DIAMONDS, RED}, {SEVEN, SPADES, BLACK}, {EIGHT, HEARTS, RED}, {NINE, DIAMONDS, RED}},
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
		t.SkipNow()
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
		// Import data into game
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

func TestNewRegister(t *testing.T) {
	r := NewRegister()
	failed := false
	failed = r.Cards == nil
	failed = r.Suits == nil
	failed = r.Ranks == nil
	failed = r.Total != 0
	if failed {
		t.Error("Register incorrectly initialized.")
	}
}

func TestAddCard(t *testing.T) {
	r := NewRegister()
	// Test all good inputs.
	for i, expected := range testCards[:52] {
		card, err := r.AddCard(testCodes[i])
		if err != nil {
			t.Errorf("Test Add %s: %v", testCodes[i], err)
		}
		if *card != *expected {
			outJSON, _ := json.Marshal(card)
			expJSON, _ := json.Marshal(expected)
			t.Errorf("Test Add %s: output did not match expected: output: %s; expected: %s", testCodes[i], outJSON, expJSON)
		}
	}
	// Test that duplicates are rejected and no more cards than total can be added.
	// TODO: consider adding different error types to ensure the correct type of error is handling the right error.
	bad := []string{"sA", "c2", "hQ", "dK", "??"}
	for _, b := range bad {
		if _, err := r.AddCard(b); err == nil {
			t.Errorf("Expected error from: (Register).AddCard(%q).", b)
		}
	}
	// Test suit constraints.
	for suit := CardSuit(0); suit < UNKNOWN_SUIT; suit++ {
		// Setup
		r = NewRegister()
		for rank := ACE; rank < UNKNOWN_RANK; rank++ {
			id := (&Card{rank, suit, 0}).Id()
			if _, err := r.AddCard(id); err != nil {
				t.Errorf("Suit constraints test setup error with %s: %v", id, err)
			}
		}
		// Test
		extra := (&Card{UNKNOWN_RANK, suit, 0}).Id()
		if _, err := r.AddCard(extra); err == nil {
			t.Errorf("Expected error from: (Register).AddCard(%q).", extra)
		}
	}
	// Test rank constraints.
	for rank := ACE; rank < UNKNOWN_RANK; rank++ {
		// Setup
		r = NewRegister()
		for suit := CardSuit(0); suit < UNKNOWN_SUIT; suit++ {
			id := (&Card{rank, suit, 0}).Id()
			if _, err := r.AddCard(id); err != nil {
				t.Errorf("Rank constraints test setup error with %s: %v", id, err)
			}
		}
		// Test
		extra := (&Card{rank, UNKNOWN_SUIT, 0}).Id()
		if _, err := r.AddCard(extra); err == nil {
			t.Errorf("Expected error from: (Register).AddCard(%q).", extra)
		}
	}
}

func TestAddCards(t *testing.T) {
	r := NewRegister()
	// Test all good inputs.
	expected := testCards[:52]
	cards, err := r.AddCards(testCodes[:52])
	if err != nil {
		t.Errorf("Test Add All: %v", err)
	}
	outJSON, omerr := json.Marshal(cards)
	expJSON, emerr := json.Marshal(expected)
	if omerr != nil || emerr != nil {
		t.Error("Testing failure:", omerr, emerr)
		t.SkipNow()
	}
	if bytes.Compare(outJSON, expJSON) != 0 {
		t.Errorf("Test Add All: output did not match expected: output: %s; expected: %s", outJSON, expJSON)
	}
	bad := []string{"sA", "c2", "hQ", "dK", "??"}
	for i := 0; i < len(bad); i++ {
		if _, err := r.AddCards(bad[i:]); err == nil {
			t.Errorf("Expected error from: (Register).AddCard(%q).", bad[i:])
		}
	}
	// Test suit constraints.
	for start := 0; start < 52; start += 13 {
		// Setup
		r = NewRegister()
		end := start + 13
		if _, err := r.AddCards(testCodes[start:end]); err != nil {
			t.Errorf("Suit of %s constraints test setup error: %v", SuitName(testCards[start].Suit), err)
		}
		// Test
		extra := (&Card{UNKNOWN_RANK, testCards[start].Suit, 0}).Id()
		if _, err := r.AddCard(extra); err == nil {
			t.Errorf("Expected error from: (Register).AddCard(%q).", extra)
		}
	}
	// Test rank constraints.
	for rank := ACE; rank < UNKNOWN_RANK; rank++ {
		// Setup
		r = NewRegister()
		ids := make([]string, 0, 4)
		for suit := CardSuit(0); suit < UNKNOWN_SUIT; suit++ {
			ids = append(ids, (&Card{rank, suit, 0}).Id())
		}
		if _, err := r.AddCards(ids); err != nil {
			t.Errorf("Rank %s constraints test setup error: %v", RankName(rank), err)
		}
		// Test
		extra := []string{(&Card{rank, UNKNOWN_SUIT, 0}).Id()}
		if _, err := r.AddCards(extra); err == nil {
			t.Errorf("Expected error from: (Register).AddCard(%q).", extra)
		}
	}
}
