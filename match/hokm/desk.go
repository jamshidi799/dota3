package hokm

import "game/model"

type desk struct {
	cards []*model.Card
}

func newDesk() *desk {
	return &desk{cards: []*model.Card{}}
}

func (d *desk) add(card *model.Card) {
	d.cards = append(d.cards, card)
}

func (d *desk) isFull() bool {
	return len(d.cards) == 4
}

func (d *desk) getCards() []*model.Card {
	return d.cards
}

func (d *desk) getSuit() model.Suit {
	return d.cards[0].Suit
}
