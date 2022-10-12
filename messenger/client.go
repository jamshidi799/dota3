package messenger

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Id         int
	Username   string
	Connection *websocket.Conn
}

func NewClient(id int, username string, connection *websocket.Conn) *Client {
	return &Client{Id: id, Username: username, Connection: connection}
}

type Clients map[int]*Client
