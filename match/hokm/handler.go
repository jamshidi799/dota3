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
	clients *messenger.Clients
	game    *game
}

func NewHandler(clients *messenger.Clients) *handler {
	return &handler{clients: clients}
}

func (h *handler) Start() error {

	h.initGame()
	h.setTrumpCaller()
	h.setTrump()
	h.dealCards()
	h.gameLoop()

	err := h.endMatch()

	// next set

	return err
}

func (h *handler) initGame() {
	players := h.getPlayers()
	h.game = newGame(players)
	h.game.start()
}

func (h *handler) getPlayers() [4]*Player {
	var players [4]*Player

	for i := 0; i < 4; i++ {
		players[i] = newPlayer((*h.clients)[i].Id, team(i%2), i, newHand(), false)
	}
	return players
}

func (h *handler) setTrumpCaller() {
	trumpCaller := rand.Intn(4)
	h.game.setTrumpCaller(trumpCaller)
	h.clients.BroadcastEvent(event.NewGameStartedEvent(trumpCaller))
}

func (h *handler) setTrump() {
	trumpCallerFiveCards := h.game.dealFirstFiveCardToTrumpCaller()
	h.sendEventToPlayer(h.game.getTrumpId(), event.NewTrumpCallerFirstCardEvent(trumpCallerFiveCards))

	var resp response.SetTrumpResponse
	h.readFromPlayer(h.game.getTrumpId(), &resp)

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

		h.sendEventToPlayer(player.id, event.NewDealCardEvent(h.game.trump, cards))
	}
}

func (h *handler) gameLoop() {
	for !h.game.isGameEnded() {
		i := 0
		for i < 4 {
			player := h.game.getCurrentPlayer()

			for retry := 0; retry < MaxRetryCount; retry++ {

				var resp response.PlayCardResponse
				h.readFromPlayer(player.id, &resp)

				card := &model.Card{
					Rank: resp.Rank,
					Suit: resp.Suit,
				}
				err := h.game.playCard(card)

				if err != nil {
					h.sendEventToPlayer(player.id, event.NewErrorEvent(err.Error()))
				} else {
					playedCardEvent := event.NewPlayedCardEvent(card, player.position)
					h.clients.BroadcastEventToOther(player.id, playedCardEvent)

					i++
					break
				}
			}

			if player == h.game.getCurrentPlayer() {
				log.Println("you should do random move here") // todo
			}
		}

		turnWinner, _ := h.game.calculateTurnResult()
		h.clients.BroadcastEvent(event.NewTurnWinnerEvent(turnWinner))
	}
}

func (h *handler) endMatch() error {
	winnerTeam, err := h.game.getWinner()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Printf("team %d won", winnerTeam)
		h.clients.BroadcastEvent(event.NewWinnerTeamEvent(h.game.score.firstTeam, h.game.score.secondTeam))
	}

	return nil
}

func (h *handler) sendEventToPlayer(playerId int, event any) {
	h.clients.SendEventToConnection(playerId, event)
}

func (h *handler) readFromPlayer(playerId int, schema any) {
	h.clients.ReadFromConnection(playerId, schema)
}
