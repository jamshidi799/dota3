package hokm

import (
	"game/messenger/dto"
	"game/model"
)

type player struct {
	model.Player
	team          team
	isTrumpCaller bool
}

func newPlayer(id int, username string, team team, position int, hand *model.Hand, isTrumpCaller bool) *player {
	return &player{
		Player:        *model.NewPlayer(id, username, position, hand),
		team:          team,
		isTrumpCaller: isTrumpCaller,
	}
}

func (p *player) toDto() dto.PlayerDto {
	return dto.PlayerDto{
		Id:            p.Id,
		Username:      p.Username,
		Team:          int(p.team),
		Position:      p.Position,
		IsTrumpCaller: p.isTrumpCaller,
	}
}
