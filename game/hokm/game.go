package hokm

import (
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

func (g *game) PlayCard(c *model.Card) {
	// check card validity

	// add card to deck
	g.desk.Add(c, g.turn)

	// check desk is full
	if g.desk.IsFull() {
		// if true: compute score
		g.turnResult()
	}

	// add turn

}

func (g *game) isCardValid(c *model.Card) {
	// get desk suit

	// if suits equal return true

	// check rad karde
}

func (g *game) turnResult() {
	// get desk cards

	// get winner pos

	// add score

	// set new leaderPos

	// refresh desk
}
