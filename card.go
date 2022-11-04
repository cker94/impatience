package main

import "log"

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
