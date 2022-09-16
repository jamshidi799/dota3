package hokm

import "game/model"

type Hand struct {
	cards map[int]model.Card // map[card int] card
}

func newHand() *Hand {
	return &Hand{
		cards: map[int]model.Card{},
	}
}

func (h *Hand) setCards(cs []model.Card) {
	cardMap := map[int]model.Card{}
	for _, card := range cs {
		cardMap[card.GetInt()] = card
	}
	h.cards = cardMap
}

func (h *Hand) appendCards(cs []model.Card) {
	for _, card := range cs {
		h.cards[card.GetInt()] = card
	}
}

func (h *Hand) popCard(i int) (*model.Card, bool) {
	card, ok := h.cards[i]
	if ok {
		delete(h.cards, i)
	}
	return &card, ok
}

func (h *Hand) hasSuit(s model.Suit) bool {
	for _, card := range h.cards {
		if card.Suit == s {
			return true
		}
	}
	return false
}

func (h *Hand) deleteCard(i int) {
	delete(h.cards, i)
}
