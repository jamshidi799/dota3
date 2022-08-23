package model

type Suit int

type Rank int

type Card struct {
	Rank Rank
	Suit Suit
}

const (
	SAPDE Suit = iota + 1
	DIAMOND
	HEART
	CLIUB
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
