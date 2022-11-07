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
	exp.Tableau.Stacks = [7][]string{
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

expected: %s`
			t.Errorf(msg, i, input, output, expected)
		}
	}
}
