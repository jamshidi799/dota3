package hokm

import (
	"game/model"
	"testing"
)

func TestSetCards(t *testing.T) {
	h := NewHand()

	cards := []model.Card{
		{1, 2},
		{14, 2},
	}
	h.SetCards(cards)

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
	d := NewHand()
	d.SetCards([]model.Card{
		{1, 2},
		{14, 2},
	})

	if !d.HasSuit(model.Suit(2)) {
		t.Fatal("hand should have suit 2")
	}

	if d.HasSuit(model.Suit(3)) {
		t.Fatal("hand should not have suit 3")
	}
}
