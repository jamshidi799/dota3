package model

type Suit int

type Rank int

type Card struct {
	Rank Rank `json:"rank"`
	Suit Suit `json:"suit"`
}

const (
	SPADE Suit = iota + 1
	DIAMOND
	HEART
	CLUB
)

const (
	Two Rank = iota + 2
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
	ACE
)

func (c *Card) GetInt() int {
	return int(c.Suit)*13 + int(c.Rank)
}
