package messenger

import (
	"encoding/json"
)

func (c *Clients) ReadFromConnection(connectionId int, schema any) {
	conn := (*c)[connectionId] // todo: remove disconnected client
	conn.read(schema)
}

func (c *Client) read(schema any) {
	_, msg, err := c.Connection.ReadMessage()
	if err != nil {
		return
	}

	err = json.Unmarshal(msg, schema)
	if err != nil {
		return
	}
}

func (c *Client) readText() string {
	_, msg, _ := c.Connection.ReadMessage()
	return string(msg)
}
