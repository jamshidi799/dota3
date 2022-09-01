package match

import (
	"fmt"
	"game/match/hokm"
	"game/model"
	"log"
)

type Match struct {
	Id          int
	Type        model.MatchType
	PlayerCount int
	Players     map[int]*model.Client
}

func (m *Match) AddPlayer(player *model.Client) {
	m.Players[player.Id] = player

	// broadcast join event to other player
	msg := []byte(fmt.Sprintf("player %d joined", player.Id))
	m.broadcastMessageToOther(player.Id, msg)

	// check if number of player is enough or not. if enough broadcast start match event
	if len(m.Players) == m.PlayerCount {
		msg := []byte("match started")
		m.broadcastMessage(msg)
	}

	m.run()
}

func (m *Match) run() {
	list := make([]*model.Client, len(m.Players))
	for i, client := range m.Players {
		list[i] = client
	}

	hokm.Run(list)
}

func (m *Match) broadcastMessage(msg []byte) {
	for _, player := range m.Players {
		if err := player.Connection.WriteMessage(1, msg); err != nil {
			log.Println(err)
		}
	}
}

func (m *Match) broadcastMessageToOther(exceptionPlayerId int, msg []byte) {
	for id, player := range m.Players {
		if id != exceptionPlayerId {
			if err := player.Connection.WriteMessage(1, msg); err != nil {
				log.Println(err)
			}
		}
	}
}

func (m *Match) sendMessageToPlayer(playerId int, msg []byte) {
	if err := m.Players[playerId].Connection.WriteMessage(1, msg); err != nil {
		log.Println(err)
	}
}
