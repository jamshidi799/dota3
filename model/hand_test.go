package model

import (
	"testing"
)

func TestSetCards(t *testing.T) {
	h := NewHand()

	cards := []Card{
		{Rank: 1, Suit: 2},
		{Rank: 14, Suit: 2},
	}
	h.SetCards(cards)

	if len(h.cards) != len(cards) {
		t.Fatalf("got %d len, wanted %d len", len(h.cards), len(cards))
	}

	for _, card := range cards {
		if _, ok := h.cards[card.GetInt()]; !ok {
			t.Fatalf("card %+v not added to Hand", card)
		}
	}

}

func TestHasSuit(t *testing.T) {
	d := NewHand()
	d.SetCards([]Card{
		{Rank: 1, Suit: 2},
		{Rank: 14, Suit: 2},
	})

	if !d.HasSuit(Suit(2)) {
		t.Fatal("Hand should have suit 2")
	}

	if d.HasSuit(Suit(3)) {
		t.Fatal("Hand should not have suit 3")
	}
}
