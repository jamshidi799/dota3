package main

func main() {

}

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//var (
//	matches = map[int]*Match{}
//)

//func main() {
//	r := gin.Default()
//
//	matches[0] = &Match{
//		1,
//		"shelem",
//		2,
//		[]Player{},
//	}
//
//	r.POST("/user/signup", signup)
//	r.POST("/match", createMatch)
//	r.GET("/join", joinMatch)
//
//	r.Run()
//}
//
//func signup(context *gin.Context) {
//}
//
//func createMatch(c *gin.Context) {
//	// return match id
//
//	matches[0] = &Match{
//		1,
//		"shelem",
//		2,
//		[]Player{},
//	}
//
//	c.JSON(200, 0)
//}
//
//func joinMatch(c *gin.Context) {
//
//	userId, _ := strconv.Atoi(c.Query("userId"))
//	matchId, _ := strconv.Atoi(c.Query("matchId"))
//
//	// create websocket connection
//
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil) // error ignored for sake of simplicity
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// add player to match
//	match := matches[matchId]
//	matches[matchId].Players = append(matches[matchId].Players, Player{userId, "ali", conn})
//
//	// broadcast join event to other player
//	for _, player := range match.Players {
//		if err := player.Connection.WriteMessage(1, []byte(fmt.Sprintf("player %d joined", userId))); err != nil {
//			log.Println(err)
//			return
//		}
//	}
//
//	// check if number of player is enough or not. if enough broadcast start match event
//	if len(match.Players) == match.PlayerCount {
//		for _, player := range match.Players {
//			if err := player.Connection.WriteMessage(1, []byte("game started")); err != nil {
//				log.Println(err)
//				return
//			}
//		}
//	}
//
//	// todo: send match info and players deck
//
//	//for {
//	//	// Read message from browser
//	//	msgType, msg, err := conn.ReadMessage()
//	//	if err != nil {
//	//		return
//	//	}
//	//
//	//	// Print the message to the console
//	//	fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
//	//
//	//	// Write message back to browser
//	//	if err = conn.WriteMessage(msgType, msg); err != nil {
//	//		return
//	//	}
//	//}
//}
