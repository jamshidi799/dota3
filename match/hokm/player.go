package hokm

import (
	"game/messenger/dto"
	"game/model"
)

type Player struct {
	id            int
	team          team
	position      int
	hand          *model.Hand
	isTrumpCaller bool
}

func newPlayer(id int, team team, position int, hand *model.Hand, isTrumpCaller bool) *Player {
	return &Player{
		id:            id,
		team:          team,
		position:      position,
		hand:          hand,
		isTrumpCaller: isTrumpCaller,
	}
}

func (p *Player) toDto() dto.PlayerDto {
	return dto.PlayerDto{
		Id:            p.id,
		Team:          int(p.team),
		Position:      p.position,
		IsTrumpCaller: p.isTrumpCaller,
	}
}
