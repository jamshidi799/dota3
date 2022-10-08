package hokm

type Player struct {
	id            int
	team          team
	position      int
	hand          *hand
	isTrumpCaller bool
}

func newPlayer(id int, team team, position int, hand *hand, isTrumpCaller bool) *Player {
	return &Player{
		id:            id,
		team:          team,
		position:      position,
		hand:          hand,
		isTrumpCaller: isTrumpCaller,
	}
}
