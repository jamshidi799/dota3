package hokm

import (
	"game/model"
	"testing"
)

func generatePlayer() [4]*Player {

	var players [4]*Player

	for i := 0; i < 4; i++ {
		players[i] = &Player{
			Team:     Team(i % 2),
			position: i,
			Hand:     NewHand(),
		}
	}
	return players
}

func TestStart(t *testing.T) {
	g := NewGame(generatePlayer())
	g.Start()

	if !g.players[g.leaderPos].IsTrumpCaller {
		t.Fatal("IsTrumpCaller not set for leader")
	}

	if g.turn != g.leaderPos {
		t.Fatalf("g.turn((%d)) != g.leaderPos((%d))", g.turn, g.leaderPos)
	}

	if len(g.players[g.leaderPos].Hand.cards) != 5 {
		t.Fatal("leader hand not set")
	}
}

func TestDealCards(t *testing.T) {
	g := NewGame(generatePlayer())
	g.Start()
	g.DealCards()

	for i, player := range g.players {
		if len(player.Hand.cards) != 13 {
			t.Fatalf("player %d hand len: %d, wanted: 13", i, len(player.Hand.cards))
		}
	}

	deck := model.NewDeck()
	for _, card := range deck.GetCards() {
		i := 0
		for _, player := range g.players {
			if _, ok := player.Hand.cards[card.GetInt()]; ok {
				i++
			}
		}

		if i != 1 {
			t.Fatalf("card %+v not found in any hand or found multiple time", card)
		}
	}

}

func TestPlayCard(t *testing.T) {
	g := NewGame(generatePlayer())
	g.Start()
	g.DealCards()

	var validCard model.Card
	for _, c := range g.players[g.turn].Hand.cards {
		validCard = c
		break
	}

	if err := g.PlayCard(&validCard); err != nil {
		t.Fatalf("error on safe move: %s", err)
	}

	if len(g.desk.cards) != 1 {
		t.Fatalf("card not added to desk")
	}

	if deskCard := g.desk.GetCards()[0]; deskCard != &validCard {
		t.Fatalf("wrong card added to desk. got %+v, want: %+v", deskCard, validCard)
	}

	if _, ok := g.players[g.leaderPos].Hand.PopCard(validCard.GetInt()); ok == true {
		t.Fatal("played card not removed from player hand")
	}

	turnBeforeInvalidMove := g.turn
	var invalidCard model.Card
	for _, c := range g.players[g.turn].Hand.cards {
		if c.Suit != g.desk.GetSuit() {
			invalidCard = c
			break
		}
	}

	if err := g.PlayCard(&invalidCard); err == nil {
		t.Fatal("not rise error on invalid move")
	}

	if g.turn != turnBeforeInvalidMove {
		t.Fatal("game turn changed after invalid move")
	}

	for _, c := range g.players[g.turn].Hand.cards {
		if c.Suit == g.desk.GetSuit() {
			validCard = c
			break
		}
	}

	if err := g.PlayCard(&validCard); err != nil {
		t.Fatalf("error on safe move: %s", err)
	}

	if len(g.desk.cards) != 2 {
		t.Fatalf("card not added to desk")
	}

}

func TestCalculateTurnResult(t *testing.T) {
	tests := []struct {
		trump     model.Suit
		desk      *Desk
		winnerPos int
	}{
		{
			model.SAPDE,
			&Desk{cards: []*model.Card{
				{12, model.DIAMOND},
				{10, model.DIAMOND},
				{14, model.DIAMOND},
				{2, model.SAPDE},
			}},
			3,
		},
		{
			model.SAPDE,
			&Desk{cards: []*model.Card{
				{2, model.SAPDE},
				{12, model.DIAMOND},
				{10, model.DIAMOND},
				{14, model.DIAMOND},
			}},
			0,
		},
		{
			model.SAPDE,
			&Desk{cards: []*model.Card{
				{12, model.DIAMOND},
				{10, model.DIAMOND},
				{14, model.DIAMOND},
				{7, model.DIAMOND},
			}},
			2,
		},
		{
			model.SAPDE,
			&Desk{cards: []*model.Card{
				{12, model.DIAMOND},
				{7, model.SAPDE},
				{14, model.DIAMOND},
				{13, model.SAPDE},
			}},
			3,
		},
	}

	for i, test := range tests {

		g := NewGame(generatePlayer())
		g.Start()

		g.trump = test.trump
		g.desk = test.desk

		leaderBeforeCalculation := g.leaderPos
		g.calculateTurnResult()

		winnerPos := (test.winnerPos + leaderBeforeCalculation) % 4
		if (g.players[winnerPos].Team == FirstTeam && g.score.FirstTeam != 1) ||
			(g.players[winnerPos].Team == SecondTeam && g.score.SecondTeam != 1) {
			t.Fatalf("test[%d]: wrong result calculation, result: %+v", i, g.score)
		}

		if g.leaderPos != winnerPos {
			t.Fatalf("test[%d]: wrong leaderPos calculation. got: %d, want: %d", i, g.leaderPos, winnerPos)
		}

		if len(g.desk.cards) != 0 {
			t.Fatalf("test[%d]: game desk is not empty after result calculation", i)
		}
	}

}
