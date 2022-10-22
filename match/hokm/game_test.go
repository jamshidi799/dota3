package hokm

import (
	"game/model"
	"testing"
)

func generatePlayer() [4]*player {

	var players [4]*player

	for i := 0; i < 4; i++ {
		players[i] = newPlayer(i, "ali", team(i%2), i, model.NewHand(), false)
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

	if len(g.players[g.leaderPos].Hand.GetCards()) != 5 {
		t.Fatal("leader Hand not set")
	}
}

func TestDealCards(t *testing.T) {
	g := newGame(generatePlayer())
	g.start()
	g.dealCards()

	for i, player := range g.players {
		if len(player.Hand.GetCards()) != 13 {
			t.Fatalf("player %d Hand len: %d, wanted: 13", i, len(player.Hand.GetCards()))
		}
	}

	deck := model.NewDeck()
	for _, card := range deck.GetCards() {
		i := 0
		for _, player := range g.players {
			if _, ok := player.Hand.GetCards()[card.GetInt()]; ok {
				i++
			}
		}

		if i != 1 {
			t.Fatalf("card %+v not found in any Hand or found multiple time", card)
		}
	}

}

func TestPlayCard(t *testing.T) {
	g := newGame(generatePlayer())
	g.start()
	g.dealCards()

	var validCard model.Card
	for _, c := range g.players[g.turn].Hand.GetCards() {
		validCard = c
		break
	}

	if err := g.playCard(&validCard); err != nil {
		t.Fatalf("error on safe move: %s", err)
	}

	if len(g.desk.GetCards()) != 1 {
		t.Fatalf("card not added to Desk")
	}

	if deskCard := g.desk.GetCards()[0]; deskCard != &validCard {
		t.Fatalf("wrong card added to Desk. got %+v, want: %+v", deskCard, validCard)
	}

	if _, ok := g.players[g.leaderPos].Hand.PopCard(validCard.GetInt()); ok == true {
		t.Fatal("played card not removed from player Hand")
	}

	turnBeforeInvalidMove := g.turn
	var invalidCard model.Card
	for _, c := range g.players[g.turn].Hand.GetCards() {
		if c.Suit != g.desk.GetSuit() {
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

	for _, c := range g.players[g.turn].Hand.GetCards() {
		if c.Suit == g.desk.GetSuit() {
			validCard = c
			break
		}
	}

	if err := g.playCard(&validCard); err != nil {
		t.Fatalf("error on safe move: %s", err)
	}

	if len(g.desk.GetCards()) != 2 {
		t.Fatalf("card not added to Desk")
	}

}

//func TestCalculateTurnResult(t *testing.T) {
//	tests := []struct {
//		trump     model.Suit
//		desk      *model.Desk
//		winnerPos int
//	}{
//		{
//			model.SAPDE,
//			&model.Desk{cards: []*model.Card{
//				{Rank: 12, Suit: model.DIAMOND},
//				{Rank: 10, Suit: model.DIAMOND},
//				{Rank: 14, Suit: model.DIAMOND},
//				{Rank: 2, Suit: model.SAPDE},
//			}},
//			3,
//		},
//		{
//			model.SAPDE,
//			&model.Desk{cards: []*model.Card{
//				{Rank: 2, Suit: model.SAPDE},
//				{Rank: 12, Suit: model.DIAMOND},
//				{Rank: 10, Suit: model.DIAMOND},
//				{Rank: 14, Suit: model.DIAMOND},
//			}},
//			0,
//		},
//		{
//			model.SAPDE,
//			&model.Desk{cards: []*model.Card{
//				{Rank: 12, Suit: model.DIAMOND},
//				{Rank: 10, Suit: model.DIAMOND},
//				{Rank: 14, Suit: model.DIAMOND},
//				{Rank: 7, Suit: model.DIAMOND},
//			}},
//			2,
//		},
//		{
//			model.SAPDE,
//			&model.Desk{cards: []*model.Card{
//				{Rank: 12, Suit: model.DIAMOND},
//				{Rank: 7, Suit: model.SAPDE},
//				{Rank: 14, Suit: model.DIAMOND},
//				{Rank: 13, Suit: model.SAPDE},
//			}},
//			3,
//		},
//	}
//
//	for i, test := range tests {
//
//		g := newGame(generatePlayer())
//		g.start()
//
//		g.trump = test.trump
//		g.desk = test.desk
//
//		leaderBeforeCalculation := g.leaderPos
//		g.calculateTurnResult()
//
//		winnerPos := (test.winnerPos + leaderBeforeCalculation) % 4
//		if (g.players[winnerPos].team == FirstTeam && g.score.firstTeam != 1) ||
//			(g.players[winnerPos].team == SecondTeam && g.score.secondTeam != 1) {
//			t.Fatalf("test[%d]: wrong result calculation, result: %+v", i, g.score)
//		}
//
//		if g.leaderPos != winnerPos {
//			t.Fatalf("test[%d]: wrong leaderPos calculation. got: %d, want: %d", i, g.leaderPos, winnerPos)
//		}
//
//		if len(g.desk.cards) != 0 {
//			t.Fatalf("test[%d]: match Desk is not empty after result calculation", i)
//		}
//	}
//
//}
