package messenger

import "github.com/gorilla/websocket"

type Client struct {
	Id         int
	Username   string
	Connection *websocket.Conn
}
