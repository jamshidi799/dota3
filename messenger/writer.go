package messenger

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Clients []*Client

// BroadcastMessage todo: remove
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

func (c Clients) BroadcastEventToOther(exceptionPlayerIndex int, event any) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	for i, client := range c {
		if i != exceptionPlayerIndex {
			if err := client.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Client) SendEventToPlayer(event any) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println(err)
	}
}
