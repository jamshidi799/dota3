package match

import (
	"fmt"
	"game/match/hokm"
	"game/messenger"
	"game/model"
)

type Match struct {
	Id      int
	Type    model.MatchType
	Clients messenger.Clients
}

func NewMatch(t model.MatchType) *Match {
	id := 1 // todo
	return &Match{Id: id, Type: t, Clients: messenger.Clients{}}
}

func (m *Match) FindClient(clientId int) *messenger.Client {
	for _, c := range m.Clients {
		if c.Id == clientId {
			return c
		}
	}
	return nil
}

func (m *Match) AddClient(client *messenger.Client) {
	if m.Type.PlayerCount <= len(m.Clients) {
		return
	}

	m.Clients[client.Id] = client

	// broadcast join event to other client
	msg := []byte(fmt.Sprintf("client %d joined", client.Id))
	m.Clients.BroadcastMessage(msg)

	if m.shouldStartMatch() {
		msg := []byte("match started")
		m.Clients.BroadcastMessage(msg)
		m.start()
	}
}

func (m *Match) shouldStartMatch() bool {
	return len(m.Clients) == m.Type.PlayerCount
}

func (m *Match) start() {
	handler := hokm.NewHandler(&m.Clients)
	handler.Start()
}
