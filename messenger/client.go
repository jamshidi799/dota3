package messenger

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Id         int
	Username   string
	Connection *websocket.Conn
}

type Clients map[int]*Client
