package match

import (
	"fmt"
	"game/match/hokm"
	"game/messenger/client"
	"game/messenger/event"
	"game/model"
)

type Match struct {
	Id       int
	Type     model.MatchType
	BotCount int
	WinScore int
	Clients  client.Clients
}

func NewMatch(id int, matchType model.MatchType, botCount int, winScore int) *Match {
	match := &Match{
		Id:       id,
		Type:     matchType,
		BotCount: botCount,
		WinScore: winScore,
		Clients:  client.Clients{},
	}

	for i := 0; i < botCount; i++ {
		match.AddClient(client.NewBotClient(i, fmt.Sprint("bot", i+1)))
	}

	return match
}

func (m *Match) FindClient(clientId int) client.Client {
	for _, c := range m.Clients {

		if c.GetId() == clientId {
			return c
		}
	}
	return nil
}

func (m *Match) AddClient(client client.Client) {
	if m.Type.PlayerCount <= len(m.Clients) {
		return
	}

	m.Clients[client.GetId()] = client

	var playersUsername []string
	for _, client := range m.Clients {
		playersUsername = append(playersUsername, client.GetUsername())
	}

	joinEvent := event.NewJoinPlayerEvent(playersUsername)
	m.Clients.BroadcastEvent(joinEvent)

	if m.shouldStartMatch() {
		m.start()
	}
}

func (m *Match) shouldStartMatch() bool {
	return len(m.Clients) == m.Type.PlayerCount
}

func (m *Match) start() {
	handler := hokm.NewHandler(m.Clients, m.WinScore)
	handler.Start()
}
