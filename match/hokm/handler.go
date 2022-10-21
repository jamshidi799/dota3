package hokm

import (
	"game/match/hokm/response"
	"game/messenger/client"
	"game/messenger/dto"
	"game/messenger/event"
	"game/model"
	"log"
	"math/rand"

	"golang.org/x/exp/maps"
)

const MaxRetryCount = 3

type handler struct {
	clients client.Clients
	game    *game
}

func NewHandler(clients client.Clients) *handler {
	return &handler{clients: clients}
}

func (h *handler) Start() error {

	h.initGame()
	h.setTrumpCaller()
	h.sendStartMatchEvent()
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

func (h *handler) getPlayers() [4]*player {
	var players [4]*player

	position := 0
	for _, client := range h.clients {
		players[position] = newPlayer(client.GetId(), team(position%2), position, model.NewHand(), false)
		position += 1
	}
	return players
}

func (h *handler) setTrumpCaller() {
	trumpCaller := rand.Intn(4)
	h.game.setTrumpCaller(trumpCaller)
}

func (h *handler) sendStartMatchEvent() {
	var players []dto.PlayerDto
	for _, player := range h.game.players {
		players = append(players, player.toDto())
	}

	h.clients.BroadcastEvent(event.NewGameStartedEvent(players))
}

func (h *handler) setTrump() {
	trumpCallerFiveCards := h.game.dealFirstFiveCardToTrumpCaller()
	h.sendEventToPlayer(h.game.getTrumpCallerId(), event.NewTrumpCallerFirstCardEvent(trumpCallerFiveCards))

	var resp response.SetTrumpResponse
	h.readFromPlayer(h.game.getTrumpCallerId(), &resp)

	if resp.Suit != 0 {
		h.game.setTrump(resp.Suit)
	}
}

func (h *handler) dealCards() {
	h.game.dealCards()

	for _, player := range h.game.players {
		hand := make(map[int]model.Card)
		maps.Copy(hand, player.Hand.GetCards())
		h.sendEventToPlayer(player.Id, event.NewDealCardEvent(h.game.trump, hand))
	}
}

func (h *handler) gameLoop() {
	for !h.game.isGameEnded() {
		i := 0
		for i < 4 {
			player := h.game.getCurrentPlayer()

			for retry := 0; retry < MaxRetryCount; retry++ {

				var resp response.PlayCardResponse
				h.readFromPlayer(player.Id, &resp)

				card := &model.Card{
					Rank: resp.Rank,
					Suit: resp.Suit,
				}
				err := h.game.playCard(card)

				if err != nil {
					h.sendEventToPlayer(player.Id, event.NewErrorEvent(err.Error()))
				} else {
					playedCardEvent := event.NewPlayedCardEvent(card, player.Position)
					h.clients.BroadcastEvent(playedCardEvent)

					i++
					break
				}
			}

			if player == h.game.getCurrentPlayer() {
				log.Println("you should do random move here") // todo
				return
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

// todo: convert any to something better
func (h *handler) sendEventToPlayer(playerId int, event any) {
	h.clients.SendEventToConnection(playerId, event)
}

func (h *handler) readFromPlayer(playerId int, schema any) {
	h.clients.ReadFromConnection(playerId, schema)
}
