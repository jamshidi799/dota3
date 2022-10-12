package model

import (
	"testing"
)

func TestAdd(t *testing.T) {
	d := NewDesk()
	card := Card{
		Rank: 4,
		Suit: 2,
	}
	d.Add(&card)

	if len(d.cards) == 0 {
		t.Fatalf("card did'nt added")
	}

	if d.cards[0] != &card {
		t.Fatalf("added card: %+v, want: %+v", d.cards[0], card)
	}
}

func TestGetSuit(t *testing.T) {
	d := NewDesk()

	suit := 2
	card := Card{
		Rank: 4,
		Suit: Suit(suit),
	}
	d.Add(&card)

	got := d.GetSuit()
	want := Suit(suit)

	if got != want {
		t.Fatalf("got: %d, wanted: %d", got, want)
	}

}
