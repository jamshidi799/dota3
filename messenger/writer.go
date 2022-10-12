package messenger

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func (c Clients) BroadcastEvent(event any) {
	for _, client := range c {
		if err := client.write(event); err != nil {
			// todo: remove disconnected client
		}
	}
}

func (c Clients) BroadcastEventToOther(exceptionPlayerId int, event any) {
	for id, client := range c {
		if id != exceptionPlayerId {
			if err := client.write(event); err != nil {
				// todo: remove disconnected client
			}
		}
	}
}

func (c *Clients) SendEventToConnection(connectionId int, event any) error {
	conn := (*c)[connectionId] // todo: remove disconnected client
	return conn.write(event)
}

func (c *Client) write(event any) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = c.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Printf("%s, id: %d", err, c.Id)
	}
	return err
}
