package hokm

import (
	"game/model"
	"log"
)

func Run() error {

	players := [4]*Player{}

	// init game
	g := NewGame(players)

	g.Start()

	// set trump
	g.SetTrump(model.DIAMOND)

	// deal
	g.DealCards()

	// play card in loop
	for g.isGameEnded() {
		for i := 0; i < 4; i++ {
			err := g.PlayCard(&model.Card{})
			if err != nil {
				return err // todo: handle error
			}
		}
	}

	// game ended
	winnerTeam, err := g.GetWinner()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Printf("team %d won", winnerTeam)
	}

	// next set

	return nil
}
