package hokm

import (
	"fmt"
	"game/messenger"
	"game/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"testing"
)

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func createHandler(t *testing.T, serverAddr string) *handler {
	var clients messenger.Clients

	for i := 0; i < 4; i++ {
		conn := mustDialWS(t, "ws"+strings.TrimPrefix(serverAddr, "http")+fmt.Sprintf("/join?matchId=0&userId=%d", i))
		client := &messenger.Client{Id: 0, Username: "ali", Connection: conn}
		clients = append(clients, client)
	}

	return NewHandler(clients)
}

func startServer() http.Handler {
	r := gin.Default()

	r.GET("/join", func(c *gin.Context) {
		_, _ = util.UpgradeConnToWebsocket(c.Writer, c.Request)
	})

	return r.Handler()
}

// todo
func TestName(t *testing.T) {
	//server := httptest.NewServer(startServer())
	//handler := createHandler(t, server.URL)
	//clients := handler.clients
	//go func() { t.Log(clients[0].ReadText()) }()
	//handler.Run()
}
