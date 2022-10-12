package model

type Player struct {
	Id       int
	Position int
	Hand     *Hand
}

func NewPlayer(id int, position int, hand *Hand) *Player {
	return &Player{Id: id, Position: position, Hand: hand}
}
