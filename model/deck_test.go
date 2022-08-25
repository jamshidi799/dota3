package model

import (
	"log"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d.cards) != 52 {
		t.Fatalf("len of card in deck is wrong")
	}

	for i, card := range d.cards {
		cardId := i + 15
		if card.GetInt() != cardId {
			log.Fatalf("not found card with id: %d in deck", cardId)
		}
	}
}

func TestDeckPop(t *testing.T) {
	d := NewDeck()
	want := d.cards[:5]

	cards := d.Pop(5)
	if !reflect.DeepEqual(want, cards) {
		t.Fatalf("want: %v but got: %v", want, cards)
	}

	if len(d.cards) != 52-5 {
		t.Fatalf("pop not remove element from deck")
	}
}
