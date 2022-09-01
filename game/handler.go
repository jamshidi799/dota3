package game

import (
	"fmt"
	"game/model"
	"log"
)

type Handler interface {
	AddPlayer(player *model.Player)
}

type BaseHandler struct {
	match *model.Match
}

func NewHandler(m *model.Match) *BaseHandler {
	return &BaseHandler{
		match: m,
	}
}

func (bh *BaseHandler) AddPlayer(player *model.Player) {
	bh.match.Players[player.Id] = player

	// broadcast join event to other player
	msg := []byte(fmt.Sprintf("player %d joined", player.Id))
	bh.broadcastMessageToOther(player.Id, msg)

	// check if number of player is enough or not. if enough broadcast start match event
	if len(bh.match.Players) == bh.match.PlayerCount {
		msg := []byte("game started")
		bh.broadcastMessage(msg)
	}

	// init game

	// todo: send game info and players deck

}

func (bh *BaseHandler) broadcastMessage(msg []byte) {
	for _, player := range bh.match.Players {
		if err := player.Connection.WriteMessage(1, msg); err != nil {
			log.Println(err)
		}
	}
}

func (bh *BaseHandler) broadcastMessageToOther(exceptionPlayerId int, msg []byte) {
	for id, player := range bh.match.Players {
		if id != exceptionPlayerId {
			if err := player.Connection.WriteMessage(1, msg); err != nil {
				log.Println(err)
			}
		}
	}
}

func (bh *BaseHandler) sendMessageToPlayer(playerId int, msg []byte) {
	if err := bh.match.Players[playerId].Connection.WriteMessage(1, msg); err != nil {
		log.Println(err)
	}
}
