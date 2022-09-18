package messenger

import (
	"encoding/json"
	"game/model"
)

func (c *Client) Read(schema any) {
	_, msg, err := c.Connection.ReadMessage()
	if err != nil {
		return
	}

	err = json.Unmarshal(msg, schema)
	if err != nil {
		return
	}
}

func (c *Client) ReadText() string {
	_, msg, _ := c.Connection.ReadMessage()
	return string(msg)
}

type TestEvent struct {
	Suit model.Suit
	Type string
}
