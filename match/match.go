package match

import (
	"fmt"
	"game/match/hokm"
	"game/messenger"
	"game/model"
)

type Match struct {
	Id          int
	Type        model.MatchType
	PlayerCount int // todo: MatchType should contain PlayerCount
	Clients     messenger.Clients
}

func NewMatch(t model.MatchType, playerCount int) *Match {
	id := 1
	return &Match{Id: id, Type: t, PlayerCount: playerCount, Clients: messenger.Clients{}}
}

func (m *Match) AddClient(client *messenger.Client) {
	m.Clients[client.Id] = client

	// broadcast join event to other client
	msg := []byte(fmt.Sprintf("client %d joined", client.Id))
	m.Clients.BroadcastMessageToOther(client.Id, msg)

	// check if number of client is enough or not. if enough broadcast start match event
	if len(m.Clients) == m.PlayerCount {
		msg := []byte("match started")
		m.Clients.BroadcastMessage(msg)
		m.run()
	}

}

func (m *Match) run() {
	hokm.Run(m.Clients)
}
