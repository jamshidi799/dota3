package hokm

import "game/model"

func Run() error {
	// init teams
	//team1 := NewTeam()
	//team2 := NewTeam()

	players := [4]*Player{}

	// init game
	g := NewGame(players)

	g.Start()

	// get trump-caller

	// set trump
	g.SetTrump(model.DIAMOND)

	// deal

	// play card in loop

	// game ended

	// next set

	return nil
}
