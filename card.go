package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type CardRank int

const (
	ACE CardRank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	UNKNOWN_RANK
)

type CardSuit int

const (
	SPADES CardSuit = iota
	CLUBS
	HEARTS
	DIAMONDS
	UNKNOWN_SUIT
)

type CardColor int

const (
	BLACK CardColor = iota
	RED
	UNKNOWN_COLOR
)

type Card struct {
	Rank  CardRank
	Suit  CardSuit
	Color CardColor
}

func (card *Card) Id() string {
	code := make([]byte, 0, 3)

	// Parse suit.
	switch card.Suit {
	case SPADES:
		code = append(code, 'S')
	case CLUBS:
		code = append(code, 'C')
	case HEARTS:
		code = append(code, 'H')
	case DIAMONDS:
		code = append(code, 'D')
	case UNKNOWN_SUIT:
		code = append(code, '?')
	default:
		log.Panicln("Out of bounds card suit:", card.Suit)
	}

	// Parse rank.
	switch card.Rank {
	case ACE:
		code = append(code, 'A')
	case TWO:
		code = append(code, '2')
	case THREE:
		code = append(code, '3')
	case FOUR:
		code = append(code, '4')
	case FIVE:
		code = append(code, '5')
	case SIX:
		code = append(code, '6')
	case SEVEN:
		code = append(code, '7')
	case EIGHT:
		code = append(code, '8')
	case NINE:
		code = append(code, '9')
	case TEN:
		code = append(code, '1', '0')
	case JACK:
		code = append(code, 'J')
	case QUEEN:
		code = append(code, 'Q')
	case KING:
		code = append(code, 'K')
	case UNKNOWN_RANK:
		code = append(code, '?')
	default:
		log.Panicln("Out of bounds card rank:", card.Rank)
	}

	return string(code)
}

func SuitName(suit CardSuit) string {
	switch suit {
	case SPADES:
		return "spades"
	case CLUBS:
		return "clubs"
	case HEARTS:
		return "hearts"
	case DIAMONDS:
		return "diamonds"
	case UNKNOWN_SUIT:
		return "unknown suit"
	}
	log.Panicln("Out of bounds card suit:", suit)
	return ""
}

func RankName(rank CardRank) string {
	switch rank {
	case ACE:
		return "ace"
	case TWO:
		return "two"
	case THREE:
		return "three"
	case FOUR:
		return "four"
	case FIVE:
		return "five"
	case SIX:
		return "six"
	case SEVEN:
		return "seven"
	case EIGHT:
		return "eight"
	case NINE:
		return "nine"
	case TEN:
		return "ten"
	case JACK:
		return "jack"
	case QUEEN:
		return "queen"
	case KING:
		return "king"
	case UNKNOWN_RANK:
		return "unknown rank"
	}
	log.Panicln("Out of bounds card rank:", rank)
	return ""
}

func ColorName(color CardColor) string {
	switch color {
	case BLACK:
		return "black"
	case RED:
		return "red"
	case UNKNOWN_COLOR:
		return "unknown color"
	}
	log.Panicln("Out of bounds card color:", color)
	return ""
}

func ParseCard(code string) (*Card, error) {
	var card Card
	size := len(code)

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
	case '?':
		card.Suit = UNKNOWN_SUIT
		card.Color = UNKNOWN_COLOR
	default:
		return nil, fmt.Errorf("Unrecognized card suit code %v in %q", suit, code)
	}

	// Get rank from second char.
	switch size {
	case 0, 1:
		return nil, errors.New("Subceeds min code length (2): " + code)
	case 2:
		switch code[1] {
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
		case '?':
			card.Rank = UNKNOWN_RANK
		default:
			return nil, fmt.Errorf("Unrecognized card rank code %v in %q.", code[1], code)
		}
	case 3:
		if code[1:3] == "10" {
			card.Rank = TEN
		} else {
			return nil, fmt.Errorf("Unrecognized card rank code %v in %q.", code[1:3], code)
		}
	default:
		return nil, errors.New("Exceeds max code length (3): " + code)
	}

	return &card, nil
}

func ParseCards(codes []string) ([]*Card, error) {
	size := len(codes)
	stack := make([]*Card, size, size)
	for i, code := range codes {
		card, err := ParseCard(code)
		if err != nil {
			return nil, err
		}
		stack[i] = card
	}
	return stack, nil
}
