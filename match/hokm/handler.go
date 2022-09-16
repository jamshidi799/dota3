package hokm

import (
	"fmt"
	"game/messenger"
	"game/messenger/event"
	"game/model"
	"log"
	"math/rand"
)

type handler struct {
	clients messenger.Clients
	game    *game
}

func NewHandler(clients messenger.Clients) *handler {
	return &handler{clients: clients}
}

func (h *handler) Run() error {

	var players [4]*Player

	for i := 0; i < 4; i++ {
		players[i] = newPlayer(h.clients[i].Id, team(i%2), i, newHand(), false, h.clients[i])
	}

	// init match
	h.game = newGame(players)
	h.game.start()

	trumpCaller := rand.Intn(4)
	h.game.setTrumpCaller(trumpCaller)
	h.clients.BroadcastEvent(event.NewGameStartedEvent(trumpCaller))

	trumpCallerFiveCards := h.game.dealFirstFiveCardToTrumpCaller()
	h.game.getTrump().client.
		SendEventToPlayer(event.NewTrumpCallerFirstCardEvent(trumpCallerFiveCards))

	// todo: listen to caller
	// set trump
	h.game.setTrump(model.DIAMOND)

	// deal
	h.dealCards()

	// play card in loop
	for !h.game.isGameEnded() {
		i := 0
		for i < 4 {
			var suit, rank int
			_, _ = fmt.Scanln(&suit, &rank)
			err := h.game.playCard(&model.Card{
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
	winnerTeam, err := h.game.getWinner()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Printf("team %d won", winnerTeam)
	}

	// next set

	return nil
}

func (h *handler) dealCards() {
	h.game.dealCards()

	for _, player := range h.game.players {
		cards := make([]model.Card, len(player.hand.cards))
		i := 0
		for _, card := range player.hand.cards {
			cards[i] = card
			i++
		}

		player.client.SendEventToPlayer(event.NewDealCardEvent(h.game.trump, cards))
	}
}
