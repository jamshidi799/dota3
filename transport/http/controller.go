package http

import (
	"fmt"
	"game/match"
	"game/messenger/client"
	"game/model"
	"game/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	matches = map[int]*match.Match{}
)

func StartServer() error {
	r := gin.Default()

	r.POST("/user/signup", signup)
	r.POST("/match", createMatch)
	r.GET("/join", joinMatch)

	matchType := model.MatchType{
		PlayerCount: 4,
		Type:        model.HOKM,
	}

	matches[0] = match.NewMatch(matchType)
	//matches[0].AddClient(client.NewBotClient(0, "ali"))
	// matches[0].AddClient(client.NewBotClient(1, "ali"))
	// matches[0].AddClient(client.NewBotClient(2, "ali"))
	// matches[0].AddClient(client.NewBotClient(3, "ali"))

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
	conn, err := util.UpgradeConnToWebsocket(&c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}

	match := matches[matchId]
	match.AddClient(client.NewUserClient(userId, "ali", conn))
	go func() {
		time.Sleep(time.Second)
		matches[0].AddClient(client.NewBotClient(3, "ali"))
		matches[0].AddClient(client.NewBotClient(20, "ali"))
		matches[0].AddClient(client.NewBotClient(3303, "ali"))
	}()
}
