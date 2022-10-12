package model

type Desk struct {
	cards []*Card
}

func NewDesk() *Desk {
	return &Desk{cards: []*Card{}}
}

func (d *Desk) Add(card *Card) {
	d.cards = append(d.cards, card)
}

func (d *Desk) IsFull() bool {
	return len(d.cards) == 4
}

func (d *Desk) GetCards() []*Card {
	return d.cards
}

func (d *Desk) GetSuit() Suit {
	return d.cards[0].Suit
}
