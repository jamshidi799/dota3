package hokm

import (
	"game/model"
	"testing"
)

func TestAdd(t *testing.T) {
	d := newDesk()
	card := model.Card{
		Rank: 4,
		Suit: 2,
	}
	d.add(&card)

	if len(d.cards) == 0 {
		t.Fatalf("card did'nt added")
	}

	if d.cards[0] != &card {
		t.Fatalf("added card: %+v, want: %+v", d.cards[0], card)
	}
}

func TestGetSuit(t *testing.T) {
	d := newDesk()

	suit := 2
	card := model.Card{
		Rank: 4,
		Suit: model.Suit(suit),
	}
	d.add(&card)

	got := d.getSuit()
	want := model.Suit(suit)

	if got != want {
		t.Fatalf("got: %d, wanted: %d", got, want)
	}

}
