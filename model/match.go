package model

import "github.com/gorilla/websocket"

type Match struct {
	Id          int
	Type        MatchType
	PlayerCount int
	Players     map[int]*Player
}

type Player struct {
	Id         int
	Username   string
	Connection *websocket.Conn
}
