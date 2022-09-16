package hokm

import (
	"errors"
	"game/model"
)

type game struct {
	deck    *model.Deck
	desk    *Desk
	players map[int]*Player

	turn      int
	leaderPos int
	trump     model.Suit
	score     score
}

func newGame(players [4]*Player) *game {
	// init deck
	deck := model.NewDeck()

	desk := newDesk()

	playerMap := map[int]*Player{}
	for _, player := range players {
		playerMap[player.position] = player
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
	g.players[g.leaderPos].hand.setCards(cards)
	return cards
}

func (g *game) getTrump() *Player {
	return g.players[g.leaderPos]
}

func (g *game) setTrump(suit model.Suit) {
	g.trump = suit
}

func (g *game) dealCards() {
	// deal remained cards
	g.players[g.leaderPos].hand.appendCards(g.deck.Pop(8))

	for i := 0; i < 4; i++ {
		if i != g.leaderPos {
			g.players[i].hand.setCards(g.deck.Pop(13))
		}
	}
}

func (g *game) playCard(c *model.Card) error {
	// check card validity
	if !g.isCardValid(c) {
		return errors.New("card is invalid")
	}

	// add card to deck
	g.desk.add(c)
	g.players[g.turn].hand.deleteCard(c.GetInt())

	// check desk is full
	if g.desk.isFull() {
		g.calculateTurnResult()
	}

	// add turn. turn start from leaderPos and go to len(players) - 1 and then restarted to 0
	g.turn = (g.turn + 1) % 4

	return nil
}

func (g *game) isCardValid(c *model.Card) bool {
	if len(g.desk.cards) == 0 {
		return true
	}

	deskSuit := g.desk.getSuit()

	if deskSuit == c.Suit {
		return true
	}

	return !g.players[g.turn].hand.hasSuit(deskSuit)
}

func (g *game) calculateTurnResult() {

	maxCard := g.desk.cards[0]
	maxIndex := 0

	for i, deskCard := range g.desk.cards[1:] {
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

	// add score
	winnerTeam := g.players[winnerPos].team
	if winnerTeam == FirstTeam {
		g.score.firstTeam += 1
	} else {
		g.score.secondTeam += 1
	}

	// set new leaderPos
	g.leaderPos = winnerPos

	// refresh desk
	g.desk = newDesk()
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

func (g *game) getCurrentPlayer() *Player {
	return g.players[g.turn]
}
