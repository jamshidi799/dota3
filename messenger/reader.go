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

type TestEvent struct {
	Suit model.Suit
	Type string
}
