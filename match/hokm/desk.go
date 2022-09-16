package hokm

import "game/model"

type Desk struct {
	cards []*model.Card
}

func newDesk() *Desk {
	return &Desk{cards: []*model.Card{}}
}

func (d *Desk) add(card *model.Card) {
	d.cards = append(d.cards, card)
}

func (d *Desk) isFull() bool {
	return len(d.cards) == 4
}

func (d *Desk) getCards() []*model.Card {
	return d.cards
}

func (d *Desk) getSuit() model.Suit {
	return d.cards[0].Suit
}
