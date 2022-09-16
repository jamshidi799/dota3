package messenger

import (
	"github.com/gorilla/websocket"
	"log"
)

type Clients []*Client

func (p Clients) BroadcastMessage(msg []byte) {
	for _, client := range p {
		if err := client.Connection.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println(err)
		}
	}
}

func (p Clients) BroadcastMessageToOther(exceptionPlayerId int, msg []byte) {
	for id, client := range p {
		if id != exceptionPlayerId {
			if err := client.Connection.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println(err)
			}
		}
	}
}

func (p Clients) SendMessageToPlayer(clientId int, msg []byte) {
	if err := p[clientId].Connection.WriteMessage(websocket.TextMessage, msg); err != nil { // fixme: use index instead of playerId
		log.Println(err)
	}
}
