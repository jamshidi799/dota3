package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeConnToWebsocket(w *gin.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upgrader.Upgrade(*w, r, nil)
}
