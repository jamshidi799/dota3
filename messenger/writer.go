package messenger

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

// BroadcastMessage todo: remove
func (c Clients) BroadcastMessage(msg []byte) {
	for _, client := range c {
		if err := client.write(msg); err != nil {
			// todo: remove disconnected client
		}
	}
}

func (c Clients) BroadcastEvent(event any) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	for _, client := range c {
		if err := client.write(data); err != nil {
			// todo: remove disconnected client
		}
	}
}

func (c Clients) BroadcastEventToOther(exceptionPlayerId int, event any) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	for id, client := range c {
		if id != exceptionPlayerId {
			if err := client.write(data); err != nil {
				// todo: remove disconnected client
			}
		}
	}
}

func (c *Clients) SendEventToConnection(connectionId int, event any) error {
	conn := (*c)[connectionId] // todo: remove disconnected client
	data, err := json.Marshal(event)
	if err != nil {
		return nil
	}

	return conn.write(data)
}

func (c *Client) write(data []byte) error {
	err := c.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Printf("%s, id: %d", err, c.Id)
	}
	return err
}
