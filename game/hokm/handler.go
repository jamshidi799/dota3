package hokm

import "game/model"

func Run() error {
	// init teams

	// init game
	g := NewGame(nil, nil)

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
