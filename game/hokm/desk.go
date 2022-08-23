package hokm

import "game/model"

type Desk struct {
	cards []*DeskCard
}

type DeskCard struct {
	card     *model.Card
	playerId int
}

func NewDesk() *Desk {
	return &Desk{cards: []*DeskCard{}}
}

func (d *Desk) Add(card *model.Card, playerId int) {
	d.cards = append(d.cards, &DeskCard{card, playerId})
}

func (d *Desk) IsFull() bool {
	return len(d.cards) == 4
}

func (d *Desk) GetCards() []*DeskCard {
	return d.cards
}

func (d *Desk) GetSuit() model.Suit {
	return d.cards[0].card.Suit
}
