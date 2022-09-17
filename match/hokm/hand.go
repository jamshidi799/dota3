package hokm

import "game/model"

type hand struct {
	cards map[int]model.Card
}

func newHand() *hand {
	return &hand{
		cards: map[int]model.Card{},
	}
}

func (h *hand) setCards(cs []model.Card) {
	cardMap := map[int]model.Card{}
	for _, card := range cs {
		cardMap[card.GetInt()] = card
	}
	h.cards = cardMap
}

func (h *hand) appendCards(cs []model.Card) {
	for _, card := range cs {
		h.cards[card.GetInt()] = card
	}
}

func (h *hand) popCard(i int) (*model.Card, bool) {
	card, ok := h.cards[i]
	if ok {
		delete(h.cards, i)
	}
	return &card, ok
}

func (h *hand) hasSuit(s model.Suit) bool {
	for _, card := range h.cards {
		if card.Suit == s {
			return true
		}
	}
	return false
}

func (h *hand) deleteCard(i int) {
	delete(h.cards, i)
}

func (h *hand) hasCard(c *model.Card) bool {
	for _, card := range h.cards {
		if card.Suit == c.Suit && card.Rank == c.Rank {
			return true
		}
	}
	return false
}
