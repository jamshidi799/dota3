package hokm

import (
	"errors"
	"game/model"
)

type game struct {
	deck    *model.Deck
	desk    *model.Desk
	players map[int]*player

	turn      int
	leaderPos int
	trump     model.Suit
	score     score
}

func newGame(players [4]*player) *game {
	// init deck
	deck := model.NewDeck()

	desk := model.NewDesk()

	playerMap := map[int]*player{}
	for _, player := range players {
		playerMap[player.Position] = player
	}

	return &game{
		deck:    deck,
		desk:    desk,
		players: playerMap,
	}
}

func (g *game) start() {
	// shuffle
	g.deck.Shuffle()
}

func (g *game) setTrumpCaller(position int) {
	g.leaderPos = position
	g.players[g.leaderPos].isTrumpCaller = true
	g.turn = g.leaderPos
}

func (g *game) dealFirstFiveCardToTrumpCaller() []model.Card {
	cards := g.deck.Pop(5)
	g.players[g.leaderPos].Hand.SetCards(cards)
	return cards
}

func (g *game) getTrumpCallerId() int {
	return g.players[g.leaderPos].Id
}

func (g *game) setTrump(suit model.Suit) {
	g.trump = suit
}

func (g *game) dealCards() {
	// deal remained cards
	g.players[g.leaderPos].Hand.AppendCards(g.deck.Pop(8))

	for i := 0; i < 4; i++ {
		if i != g.leaderPos {
			g.players[i].Hand.SetCards(g.deck.Pop(13))
		}
	}
}

func (g *game) playCard(c *model.Card) error {
	// check card validity
	if err := g.isCardValid(c); err != nil {
		return err
	}

	// Add card to deck
	g.desk.Add(c)
	g.players[g.turn].Hand.DeleteCard(c.GetInt())

	// Add turn. turn start from leaderPos and go to len(players) - 1 and then restarted to 0
	g.turn = (g.turn + 1) % 4

	return nil
}

func (g *game) isCardValid(c *model.Card) error {

	if !g.players[g.turn].Hand.HasCard(c) {
		return errors.New("invalid card. card not found in you Hand")
	}

	if len(g.desk.GetCards()) == 0 {
		return nil
	}

	deskSuit := g.desk.GetSuit()

	if deskSuit == c.Suit {
		return nil
	}

	if g.players[g.turn].Hand.HasSuit(deskSuit) {
		return errors.New("card with invalid suit")
	}

	return nil
}

func (g *game) calculateTurnResult() (int, error) {
	if !g.desk.IsFull() {
		return -1, errors.New("turn is not ended")
	}

	maxCard := g.desk.GetCards()[0]
	maxIndex := 0

	for i, deskCard := range g.desk.GetCards()[1:] {
		if maxCard.Suit == deskCard.Suit && maxCard.Rank < deskCard.Rank {
			maxCard = deskCard
			maxIndex = i + 1
		} else if deskCard.Suit == g.trump {
			maxCard = deskCard
			maxIndex = i + 1
		}
	}

	// get winner pos
	winnerPos := (maxIndex + g.leaderPos) % 4

	// Add score
	winnerTeam := g.players[winnerPos].team
	if winnerTeam == FirstTeam {
		g.score.firstTeam += 1
	} else {
		g.score.secondTeam += 1
	}

	// set new leaderPos
	g.leaderPos = winnerPos
	g.turn = winnerPos

	// refresh Desk
	g.desk = model.NewDesk()

	return winnerPos, nil
}

func (g *game) isGameEnded() bool {
	return g.score.firstTeam == 7 || g.score.secondTeam == 7
}

func (g *game) getWinner() (team, error) {
	if g.score.firstTeam == 7 {
		return FirstTeam, nil
	} else if g.score.secondTeam == 7 {
		return SecondTeam, nil
	}

	return 0, errors.New("match not finished yet")
}

func (g *game) getCurrentPlayer() *player {
	return g.players[g.turn]
}
