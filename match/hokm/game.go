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
	g.players[g.leaderPos].IsTrumpCaller = true
	g.turn = g.leaderPos

	// deal first 5 card to trump-caller
	g.players[g.leaderPos].Hand.SetCards(g.deck.Pop(5))
}

func (g *game) GetTrump() *Player {
	return g.players[g.leaderPos]
}

func (g *game) SetTrump(suit model.Suit) {
	g.trump = suit
}

func (g *game) DealCards() {
	// deal remained cards
	g.players[g.leaderPos].Hand.AppendCards(g.deck.Pop(8))

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
	g.desk.Add(c)
	g.players[g.turn].Hand.DeleteCard(c.GetInt())

	// check desk is full
	if g.desk.IsFull() {
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

	deskSuit := g.desk.GetSuit()

	if deskSuit == c.Suit {
		return true
	}

	return !g.players[g.turn].Hand.HasSuit(deskSuit)
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
	winnerTeam := g.players[winnerPos].Team
	if winnerTeam == FirstTeam {
		g.score.FirstTeam += 1
	} else {
		g.score.SecondTeam += 1
	}

	// set new leaderPos
	g.leaderPos = winnerPos

	// refresh desk
	g.desk = NewDesk()
}

func (g *game) isGameEnded() bool {
	return g.score.FirstTeam == 7 || g.score.SecondTeam == 7
}

func (g *game) GetWinner() (Team, error) {
	if g.score.FirstTeam == 7 {
		return FirstTeam, nil
	} else if g.score.SecondTeam == 7 {
		return SecondTeam, nil
	}

	return 0, errors.New("match not finished yet")
}
