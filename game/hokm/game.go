package hokm

import "game/model"

type game struct {
	deck  *model.Deck
	team1 *Team
	team2 *Team
	desk  *Desk

	turn      int
	leaderPos int
	trump     model.Suit
}

func NewGame(team1, team2 *Team) *game {
	// init deck
	deck := model.NewDeck()

	desk := NewDesk()

	return &game{
		deck:  deck,
		team1: team1,
		team2: team2,
		desk:  desk,
	}
}

func (g *game) Start() {
	// shuffle

	// set trump-caller

	// deal first 5 card to trump-caller
}

func (g *game) SetTrump(suit model.Suit) {

}

func (g *game) DealCards() {
	// deal remained cards
}

func (g *game) PlayCard(c *model.Card) {
	// check card validity

	// add card to deck
	g.desk.Add(c)

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
}
