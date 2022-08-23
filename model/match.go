package model

type Match struct {
	Id          int
	Type        string
	PlayerCount int
	Players     []Player
}

type Player struct {
	Id       int
	Username string
}
