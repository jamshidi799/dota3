package hokm

import (
	"game/messenger"
)

type Player struct {
	Id            int
	Team          Team
	position      int
	Hand          *Hand
	IsTrumpCaller bool
	Client        *messenger.Client
}

func NewPlayer(id int, team Team, position int, hand *Hand, isTrumpCaller bool, client *messenger.Client) *Player {
	return &Player{
		Id:            id,
		Team:          team,
		position:      position,
		Hand:          hand,
		IsTrumpCaller: isTrumpCaller,
		Client:        client,
	}
}
