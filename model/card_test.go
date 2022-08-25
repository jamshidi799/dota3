package model

import "testing"

func TestGetInt(t *testing.T) {
	want := 52
	KingHeartCard := Card{
		Rank: 13,
		Suit: 3,
	}

	if KingHeartCard.GetInt() != want {
		t.Fatalf("got %d, wanted: %d", KingHeartCard.GetInt(), want)
	}
}
