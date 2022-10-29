package impatience

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type savedGame struct {
	Stock struct {
		List      []string
		Pos       uint8
		LoopCount uint8
	}
	Tableau     [7][]string
	Foundations [4][]string
}

// TODO: Load game from file
func LoadGame(path string) (*Game, error) {
	var save savedGame
	// Open file.
	if file, err := os.Open(path); err != nil {
		return nil, err
	} else {
		// Whatever happens, close file when function exits.
		defer file.Close()
		// Read file into memory.
		if contents, err := io.ReadAll(file); err != nil {
			return nil, err
		} else {
			// Check if format is JSON or TOML.
			var err error
			ext := filepath.Ext(path)
			switch {
			case ext == ".json":
				json.Unmarshal(contents, &save)
			case ext == ".toml":
				toml.Unmarshal(contents, &save)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	game := NewGame()

	// Load stock from save data.
	game.StockPos = save.Stock.Pos
	game.StockLoop = save.Stock.LoopCount
	size := len(save.Stock.List)
	stockList := make([]*Card, size, size)
	for i, code := range save.Stock.List {
		stockList[i] = ParseCardCode(code)
	}
	game.Stacks[STOCK] = stockList

	return game, nil
}

func ParseCardCode(code string) *Card {
	var card Card
	// Normalize casing to make parsing case-insensitve.
	code = strings.ToUpper(code)

	// Get suit from first char.
	suit := code[0]
	switch suit {
	case 'S':
		card.Suit = SPADES
		card.Color = BLACK
	case 'C':
		card.Suit = CLUBS
		card.Color = BLACK
	case 'H':
		card.Suit = HEARTS
		card.Color = RED
	case 'D':
		card.Suit = DIAMONDS
		card.Color = RED
	default:
		card.Suit = UNKNOWN_SUIT
		card.Color = UNKNOWN_COLOR
	}

	// Get rank from second char
	rank := code[1]
	switch rank {
	case 'A':
		card.Rank = ACE
	case '2':
		card.Rank = TWO
	case '3':
		card.Rank = THREE
	case '4':
		card.Rank = FOUR
	case '5':
		card.Rank = FIVE
	case '6':
		card.Rank = SIX
	case '7':
		card.Rank = SEVEN
	case '8':
		card.Rank = EIGHT
	case '9':
		card.Rank = NINE
	case '1':
		card.Rank = TEN
	case 'J':
		card.Rank = JACK
	case 'Q':
		card.Rank = QUEEN
	case 'K':
		card.Rank = KING
	default:
		card.Rank = UNKNOWN_RANK
	}

	// Determine if card is faceup or facedown from last char.
	facing := code[len(code)-1]
	card.Faceup = facing != 'D'

	return &card
}
