package messenger

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Clients []*Client

func (c Clients) BroadcastMessage(msg []byte) {
	for _, client := range c {
		if err := client.Connection.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println(err)
		}
	}
}

func (c Clients) BroadcastEvent(event any) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	for _, client := range c {
		if err := client.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println(err)
		}
	}
}

func (c Clients) BroadcastMessageToOther(exceptionPlayerId int, msg []byte) {
	for id, client := range c {
		if id != exceptionPlayerId {
			if err := client.Connection.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println(err)
			}
		}
	}
}

func (c Clients) SendMessageToPlayer(clientId int, msg []byte) {
	if err := c[clientId].Connection.WriteMessage(websocket.TextMessage, msg); err != nil { // fixme: use index instead of playerId
		log.Println(err)
	}
}

func (c Client) SendMessageToPlayer(msg []byte) {
	if err := c.Connection.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println(err)
	}
}

func (c Client) SendEventToPlayer(event any) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println(err)
	}
}
