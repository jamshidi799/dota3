package http

import (
	"fmt"
	"game/match"
	"game/messenger"
	"game/model"
	"game/util"
	"github.com/gin-gonic/gin"
	"strconv"
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
	conn, err := util.UpgradeConnToWebsocket(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}

	match := matches[matchId]
	match.AddClient(messenger.NewClient(userId, "ali", conn))
}
