package hokm

import "game/model"

type Hand struct {
	cards map[int]model.Card // map[card int] card
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) SetCards(cs []model.Card) {
}

func (h *Hand) PopCard(i int) (*model.Card, bool) {
	return nil, true
}

func (h *Hand) HasSuit(s model.Suit) bool {
	// todo
	return true
}
