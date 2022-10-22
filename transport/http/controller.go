package http

import (
	"fmt"
	"game/match"
	"game/messenger/client"
	"game/model"
	"game/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	matchId := int(uuid.New().ID())
	matchType := model.MatchType{
		PlayerCount: 4,
		Type:        model.HOKM,
	}

	matches[matchId] = match.NewMatch(matchType)
	c.JSON(http.StatusOK, gin.H{
		"id": matchId,
	})
}

func joinMatch(c *gin.Context) {
	username := c.Query("username")
	matchId, _ := strconv.Atoi(c.Query("matchId"))

	conn, err := util.UpgradeConnToWebsocket(&c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}

	match := matches[matchId]

	match.AddClient(client.NewUserClient(len(match.Clients), username, conn))
	go func() {
		time.Sleep(time.Second)
		matches[0].AddClient(client.NewBotClient(1, "bot1"))
		matches[0].AddClient(client.NewBotClient(3, "bot3"))
		matches[0].AddClient(client.NewBotClient(2, "bot2"))
	}()
}
