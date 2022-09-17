package hokm

import (
	"game/messenger"
)

type Player struct {
	id            int
	team          team
	position      int
	hand          *hand
	isTrumpCaller bool
	client        *messenger.Client
}

func newPlayer(id int, team team, position int, hand *hand, isTrumpCaller bool, client *messenger.Client) *Player {
	return &Player{
		id:            id,
		team:          team,
		position:      position,
		hand:          hand,
		isTrumpCaller: isTrumpCaller,
		client:        client,
	}
}
