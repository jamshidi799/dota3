package hokm

import (
	"fmt"
	"game/messenger"
	"game/messenger/event"
	"game/model"
	"log"
	"math/rand"
)

func Run(clients messenger.Clients) error {

	var players [4]*Player

	for i := 0; i < 4; i++ {
		players[i] = NewPlayer(clients[i].Id, Team(i%2), i, NewHand(), false, clients[i])
	}

	// init match
	g := NewGame(players)

	g.Start()

	trumpCaller := rand.Intn(4)
	g.setTrumpCaller(trumpCaller)
	clients.BroadcastEvent(event.NewGameStartedEvent(trumpCaller))

	trumpCallerFiveCards := g.dealFirstFiveCardToTrumpCaller()
	g.GetTrump().Client.
		SendEventToPlayer(event.NewTrumpCallerFirstCardEvent(trumpCallerFiveCards))

	// todo: listen to caller
	// set trump
	g.SetTrump(model.DIAMOND)

	// deal
	g.DealCards()

	// play card in loop
	for !g.isGameEnded() {
		i := 0
		for i < 4 {
			var suit, rank int
			_, _ = fmt.Scanln(&suit, &rank)
			err := g.PlayCard(&model.Card{
				Rank: model.Rank(rank),
				Suit: model.Suit(suit),
			})
			if err != nil {
				log.Printf(err.Error())
			} else {
				i++
			}
		}
	}

	// match ended
	winnerTeam, err := g.GetWinner()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Printf("team %d won", winnerTeam)
	}

	// next set

	return nil
}
