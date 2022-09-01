package http

import (
	"fmt"
	"game/game"
	"game/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var (
	handlers = map[int]*game.BaseHandler{}
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

	return r.Run()
}

func signup(context *gin.Context) {
}

func createMatch(c *gin.Context) {
	// return match id

	handlers[0] = game.NewHandler(&model.Match{
		Id:          1,
		Type:        model.HOKM4,
		PlayerCount: 4,
		Players:     map[int]*model.Player{},
	})

	c.JSON(200, 0)
}

func joinMatch(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Query("userId"))
	matchId, _ := strconv.Atoi(c.Query("matchId"))

	// create websocket connection

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil) // error ignored for sake of simplicity
	if err != nil {
		fmt.Println(err)
		return
	}

	// add player to handler
	handler := handlers[matchId]
	handler.AddPlayer(&model.Player{Id: userId, Username: "ali", Connection: conn})

	//for {
	//	// Read message from browser
	//	msgType, msg, err := conn.ReadMessage()
	//	if err != nil {
	//		return
	//	}
	//
	//	// Print the message to the console
	//	fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
	//
	//	// Write message back to browser
	//	if err = conn.WriteMessage(msgType, msg); err != nil {
	//		return
	//	}
	//}
}
