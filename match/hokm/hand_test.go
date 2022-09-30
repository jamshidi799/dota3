package hokm

import (
	"game/model"
	"testing"
)

func TestSetCards(t *testing.T) {
	h := newHand()

	cards := []model.Card{
		{Rank: 1, Suit: 2},
		{Rank: 14, Suit: 2},
	}
	h.setCards(cards)

	if len(h.cards) != len(cards) {
		t.Fatalf("got %d len, wanted %d len", len(h.cards), len(cards))
	}

	for _, card := range cards {
		if _, ok := h.cards[card.GetInt()]; !ok {
			t.Fatalf("card %+v not added to hand", card)
		}
	}

}

func TestHasSuit(t *testing.T) {
	d := newHand()
	d.setCards([]model.Card{
		{Rank: 1, Suit: 2},
		{Rank: 14, Suit: 2},
	})

	if !d.hasSuit(model.Suit(2)) {
		t.Fatal("hand should have suit 2")
	}

	if d.hasSuit(model.Suit(3)) {
		t.Fatal("hand should not have suit 3")
	}
}
