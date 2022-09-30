package hokm

import (
	"game/model"
	"testing"
)

func generatePlayer() [4]*Player {

	var players [4]*Player

	for i := 0; i < 4; i++ {
		players[i] = &Player{
			team:     team(i % 2),
			position: i,
			hand:     newHand(),
		}
	}
	return players
}

func TestStart(t *testing.T) {
	g := newGame(generatePlayer())
	g.start()

	if !g.players[g.leaderPos].isTrumpCaller {
		t.Fatal("isTrumpCaller not set for leader")
	}

	if g.turn != g.leaderPos {
		t.Fatalf("g.turn((%d)) != g.leaderPos((%d))", g.turn, g.leaderPos)
	}

	if len(g.players[g.leaderPos].hand.cards) != 5 {
		t.Fatal("leader hand not set")
	}
}

func TestDealCards(t *testing.T) {
	g := newGame(generatePlayer())
	g.start()
	g.dealCards()

	for i, player := range g.players {
		if len(player.hand.cards) != 13 {
			t.Fatalf("player %d hand len: %d, wanted: 13", i, len(player.hand.cards))
		}
	}

	deck := model.NewDeck()
	for _, card := range deck.GetCards() {
		i := 0
		for _, player := range g.players {
			if _, ok := player.hand.cards[card.GetInt()]; ok {
				i++
			}
		}

		if i != 1 {
			t.Fatalf("card %+v not found in any hand or found multiple time", card)
		}
	}

}

func TestPlayCard(t *testing.T) {
	g := newGame(generatePlayer())
	g.start()
	g.dealCards()

	var validCard model.Card
	for _, c := range g.players[g.turn].hand.cards {
		validCard = c
		break
	}

	if err := g.playCard(&validCard); err != nil {
		t.Fatalf("error on safe move: %s", err)
	}

	if len(g.desk.cards) != 1 {
		t.Fatalf("card not added to desk")
	}

	if deskCard := g.desk.getCards()[0]; deskCard != &validCard {
		t.Fatalf("wrong card added to desk. got %+v, want: %+v", deskCard, validCard)
	}

	if _, ok := g.players[g.leaderPos].hand.popCard(validCard.GetInt()); ok == true {
		t.Fatal("played card not removed from player hand")
	}

	turnBeforeInvalidMove := g.turn
	var invalidCard model.Card
	for _, c := range g.players[g.turn].hand.cards {
		if c.Suit != g.desk.getSuit() {
			invalidCard = c
			break
		}
	}

	if err := g.playCard(&invalidCard); err == nil {
		t.Fatal("not rise error on invalid move")
	}

	if g.turn != turnBeforeInvalidMove {
		t.Fatal("match turn changed after invalid move")
	}

	for _, c := range g.players[g.turn].hand.cards {
		if c.Suit == g.desk.getSuit() {
			validCard = c
			break
		}
	}

	if err := g.playCard(&validCard); err != nil {
		t.Fatalf("error on safe move: %s", err)
	}

	if len(g.desk.cards) != 2 {
		t.Fatalf("card not added to desk")
	}

}

func TestCalculateTurnResult(t *testing.T) {
	tests := []struct {
		trump     model.Suit
		desk      *desk
		winnerPos int
	}{
		{
			model.SAPDE,
			&desk{cards: []*model.Card{
				{Rank: 12, Suit: model.DIAMOND},
				{Rank: 10, Suit: model.DIAMOND},
				{Rank: 14, Suit: model.DIAMOND},
				{Rank: 2, Suit: model.SAPDE},
			}},
			3,
		},
		{
			model.SAPDE,
			&desk{cards: []*model.Card{
				{Rank: 2, Suit: model.SAPDE},
				{Rank: 12, Suit: model.DIAMOND},
				{Rank: 10, Suit: model.DIAMOND},
				{Rank: 14, Suit: model.DIAMOND},
			}},
			0,
		},
		{
			model.SAPDE,
			&desk{cards: []*model.Card{
				{Rank: 12, Suit: model.DIAMOND},
				{Rank: 10, Suit: model.DIAMOND},
				{Rank: 14, Suit: model.DIAMOND},
				{Rank: 7, Suit: model.DIAMOND},
			}},
			2,
		},
		{
			model.SAPDE,
			&desk{cards: []*model.Card{
				{Rank: 12, Suit: model.DIAMOND},
				{Rank: 7, Suit: model.SAPDE},
				{Rank: 14, Suit: model.DIAMOND},
				{Rank: 13, Suit: model.SAPDE},
			}},
			3,
		},
	}

	for i, test := range tests {

		g := newGame(generatePlayer())
		g.start()

		g.trump = test.trump
		g.desk = test.desk

		leaderBeforeCalculation := g.leaderPos
		g.calculateTurnResult()

		winnerPos := (test.winnerPos + leaderBeforeCalculation) % 4
		if (g.players[winnerPos].team == FirstTeam && g.score.firstTeam != 1) ||
			(g.players[winnerPos].team == SecondTeam && g.score.secondTeam != 1) {
			t.Fatalf("test[%d]: wrong result calculation, result: %+v", i, g.score)
		}

		if g.leaderPos != winnerPos {
			t.Fatalf("test[%d]: wrong leaderPos calculation. got: %d, want: %d", i, g.leaderPos, winnerPos)
		}

		if len(g.desk.cards) != 0 {
			t.Fatalf("test[%d]: match desk is not empty after result calculation", i)
		}
	}

}
