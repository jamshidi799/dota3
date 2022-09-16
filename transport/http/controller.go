package http

import (
	"fmt"
	"game/match"
	"game/messenger"
	"game/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var (
	matches  = map[int]*match.Match{}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func StartServer() error {
	r := gin.Default()

	r.POST("/user/signup", signup)
	r.POST("/match", createMatch)
	r.GET("/join", joinMatch)

	matches[0] = match.NewMatch(model.HOKM4, 4)

	return r.Run()
}

func signup(c *gin.Context) {
	// todo
}

func createMatch(c *gin.Context) {
	// todo
}

func joinMatch(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Query("userId"))
	matchId, _ := strconv.Atoi(c.Query("matchId"))

	// create websocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	handler := matches[matchId]
	handler.AddClient(&messenger.Client{Id: userId, Username: "ali", Connection: conn})
}
