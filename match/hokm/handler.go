package hokm

import (
	"game/match/hokm/response"
	"game/messenger"
	"game/messenger/event"
	"game/model"
	"log"
	"math/rand"
)

const MaxRetryCount = 3

type handler struct {
	clients messenger.Clients
	game    *game
}

func NewHandler(clients messenger.Clients) *handler {
	return &handler{clients: clients}
}

func (h *handler) Run() error {

	h.initGame()
	h.setTrumpCaller()
	h.setTrump()
	h.dealCards()
	h.gameLoop()

	// match ended
	winnerTeam, err := h.game.getWinner()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Printf("team %d won", winnerTeam)
		h.clients.BroadcastEvent(event.NewWinnerTeamEvent(h.game.score.firstTeam, h.game.score.secondTeam))
	}

	// next set

	return nil
}

func (h *handler) initGame() {
	var players [4]*Player

	for i := 0; i < 4; i++ {
		players[i] = newPlayer(h.clients[i].Id, team(i%2), i, newHand(), false, h.clients[i])
	}

	h.game = newGame(players)
	h.game.start()
}

func (h *handler) setTrumpCaller() {
	trumpCaller := rand.Intn(4)
	h.game.setTrumpCaller(trumpCaller)
	h.clients.BroadcastEvent(event.NewGameStartedEvent(trumpCaller))
}

func (h *handler) setTrump() {
	trumpCallerFiveCards := h.game.dealFirstFiveCardToTrumpCaller()
	h.game.getTrump().client.
		SendEventToPlayer(event.NewTrumpCallerFirstCardEvent(trumpCallerFiveCards))

	var resp response.SetTrumpResponse
	h.game.getTrump().client.Read(&resp)
	if resp.Suit != 0 {
		h.game.setTrump(resp.Suit)
	}
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

func (h *handler) gameLoop() {
	for !h.game.isGameEnded() {
		i := 0
		for i < 4 {
			player := h.game.getCurrentPlayer()

			for retry := 0; retry < MaxRetryCount; retry++ {

				var resp response.PlayCardResponse
				player.client.Read(&resp)

				card := &model.Card{
					Rank: resp.Rank,
					Suit: resp.Suit,
				}
				err := h.game.playCard(card)

				if err != nil {
					player.client.SendEventToPlayer(event.NewErrorEvent(err.Error()))
				} else {
					playedCardEvent := event.NewPlayedCardEvent(card, player.position)
					h.clients.BroadcastEventToOther(player.position, playedCardEvent)

					i++
					break
				}
			}
		}

		turnWinner, _ := h.game.calculateTurnResult()
		h.clients.BroadcastEvent(event.NewTurnWinnerEvent(turnWinner))
	}
}
