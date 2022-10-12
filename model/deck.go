package model

import (
	"math/rand"
	"time"
)

// Deck todo: use []*Card
type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	cards := make([]Card, 52)
	index := 0
	for i := 1; i < 5; i++ {
		for j := 2; j < 15; j++ {
			cards[index] = Card{
				Rank: Rank(j),
				Suit: Suit(i),
			}
			index++
		}
	}

	return &Deck{cards: cards}
}

func (d *Deck) GetCards() []Card {
	return d.cards
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *Deck) Pop(count int) []Card {
	cards := d.cards[:count]
	if len(d.cards) > 1 {
		d.cards = d.cards[count:]
	} else {
		d.cards = make([]Card, 0)
	}
	return cards
}
