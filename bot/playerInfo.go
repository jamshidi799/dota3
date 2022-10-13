package bot

import "game/messenger/dto"

type playerInfo struct {
	Id            int
	Team          int
	Position      int
	IsTrumpCaller bool
}

func fromPlayerDto(dto *dto.PlayerDto) *playerInfo {
	return &playerInfo{
		Id:            dto.Id,
		Team:          dto.Team,
		Position:      dto.Position,
		IsTrumpCaller: dto.IsTrumpCaller,
	}
}
