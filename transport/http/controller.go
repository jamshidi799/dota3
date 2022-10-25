package http

import (
	"fmt"
	"game/match"
	"game/messenger/client"
	"game/model"
	"game/transport/http/request"
	"game/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	matches = map[int]*match.Match{}
)

func StartServer() error {
	r := gin.Default()

	r.POST("/match", createMatch)
	r.GET("/join", joinMatch)

	return r.Run()
}

func createMatch(c *gin.Context) {
	request := request.CreateMatchRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	matchId := int(uuid.New().ID()) % 100
	matchType := model.MatchType{
		PlayerCount: 4,
		Type:        model.HOKM,
	}

	matches[matchId] = match.NewMatch(matchId, matchType, request.BotCount, request.WinScore)

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
}
