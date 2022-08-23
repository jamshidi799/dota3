package hokm

import (
	"errors"
	"game/model"
	"math/rand"
)

type game struct {
	deck    *model.Deck
	desk    *Desk
	players map[int]*Player

	turn      int
	leaderPos int
	trump     model.Suit
	score     Score
}

func NewGame(players [4]*Player) *game {
	// init deck
	deck := model.NewDeck()

	desk := NewDesk()

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

func (g *game) Start() {
	// shuffle
	g.deck.Shuffle()

	// set trump-caller
	g.leaderPos = rand.Intn(4)

	// deal first 5 card to trump-caller
	g.players[g.leaderPos].Hand.SetCards(g.deck.Pop(5))
}

func (g *game) SetTrump(suit model.Suit) {
	g.trump = suit
}

func (g *game) DealCards() {
	// deal remained cards
	g.players[g.leaderPos].Hand.SetCards(g.deck.Pop(8))

	for i := 0; i < 4; i++ {
		if i != g.leaderPos {
			g.players[i].Hand.SetCards(g.deck.Pop(13))
		}
	}
}

func (g *game) PlayCard(c *model.Card) error {
	// check card validity
	if !g.isCardValid(c) {
		return errors.New("card is invalid")
	}

	// add card to deck
	g.desk.Add(c, g.turn)

	// check desk is full
	if g.desk.IsFull() {
		// if true: compute score
		g.turnResult()
	}

	// add turn. turn start from leaderPos and go to len(players) - 1 and then restarted to 0

	return nil
}

func (g *game) isCardValid(c *model.Card) bool {
	deskSuit := g.desk.GetSuit()

	if deskSuit == c.Suit {
		return true
	}

	return !g.players[g.turn].Hand.HasSuit(deskSuit)
}

func (g *game) turnResult() {
	// get desk cards

	// get winner pos

	// add score

	// set new leaderPos

	// refresh desk
}
