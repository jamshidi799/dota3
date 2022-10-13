package model

type Hand struct {
	cards map[int]Card
}

func NewHandFromCards(cards map[int]Card) *Hand {
	return &Hand{cards: cards}
}

func NewHand() *Hand {
	return &Hand{
		cards: map[int]Card{},
	}
}

func (h *Hand) GetCards() map[int]Card {
	return h.cards
}

func (h *Hand) SetCards(cs []Card) {
	cardMap := map[int]Card{}
	for _, card := range cs {
		cardMap[card.GetInt()] = card
	}
	h.cards = cardMap
}

func (h *Hand) AppendCards(cs []Card) {
	for _, card := range cs {
		h.cards[card.GetInt()] = card
	}
}

func (h *Hand) PopCard(i int) (*Card, bool) {
	card, ok := h.cards[i]
	if ok {
		delete(h.cards, i)
	}
	return &card, ok
}

func (h *Hand) HasSuit(s Suit) bool {
	for _, card := range h.cards {
		if card.Suit == s {
			return true
		}
	}
	return false
}

func (h *Hand) DeleteCard(i int) {
	delete(h.cards, i)
}

func (h *Hand) HasCard(c *Card) bool {
	for _, card := range h.cards {
		if card.Suit == c.Suit && card.Rank == c.Rank {
			return true
		}
	}
	return false
}
