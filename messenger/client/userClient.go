package client

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type UserClient struct {
	BaseClient
	Connection *websocket.Conn
}

func NewUserClient(id int, username string, connection *websocket.Conn) *UserClient {
	return &UserClient{
		BaseClient: BaseClient{
			id:       id,
			username: username,
		},
		Connection: connection,
	}
}

func (u *UserClient) GetId() int {
	return u.id
}

func (u *UserClient) GetUsername() string {
	return u.username
}

func (u *UserClient) send(event any) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = u.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Printf("%s, id: %d", err, u.id)
	}
	return err
}

func (u *UserClient) receive(schema any) {
	_, msg, err := u.Connection.ReadMessage()
	if err != nil {
		return
	}

	err = json.Unmarshal(msg, schema)
	if err != nil {
		return
	}
}
