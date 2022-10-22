package model

type Player struct {
	Id       int
	Username string
	Position int
	Hand     *Hand
}

func NewPlayer(id int, username string, position int, hand *Hand) *Player {
	return &Player{Id: id, Username: username, Position: position, Hand: hand}
}
