package hokm

import "game/model"

type Desk struct {
	cards []*model.Card
}

func NewDesk() *Desk {
	return &Desk{cards: []*model.Card{}}
}

func (d *Desk) Add(card *model.Card) {
	d.cards = append(d.cards, card)
}

func (d *Desk) IsFull() bool {
	return len(d.cards) == 4
}

func (d *Desk) GetCards() []*model.Card {
	return d.cards
}

func (d *Desk) GetSuit() model.Suit {
	return d.cards[0].Suit
}
