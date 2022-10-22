package bot

import (
	"game/match/hokm/response"
	"game/messenger/dto"
	"game/messenger/event"
	"game/model"

	"golang.org/x/exp/slices"
)

type HokmBot struct {
	id          int
	info        *playerInfo
	teammate    *playerInfo
	trump       model.Suit
	desk        []*model.Card
	hand        map[int]model.Card
	playedCards map[int]bool
}

func NewHokmBot(id int) *HokmBot {
	return &HokmBot{
		id: id,
	}
}

func (h *HokmBot) Receive(msg any) error {
	switch msg.(type) {

	case *event.GameStartedEvent:
		h.handleGameStartedEvent(msg.(*event.GameStartedEvent))

	case *event.TrumpCallerFirstCardEvent:
		h.handleTrumpCallerFirstCardEvent(msg.(*event.TrumpCallerFirstCardEvent))

	case *event.DealCardEvent:
		h.handleDealCardEvent(msg.(*event.DealCardEvent))

	case *event.PlayedCardEvent:
		h.handlePlayedCardEvent(msg.(*event.PlayedCardEvent))

	case *event.TurnWinnerEvent:
		h.handleTurnWinnerEvent(msg.(*event.TurnWinnerEvent))
	}

	return nil
}

func (h *HokmBot) handleGameStartedEvent(e *event.GameStartedEvent) {
	h.setBotInfo(e.Players)
	h.setTeammateInfo(e.Players)
}

func (h *HokmBot) setBotInfo(players []dto.PlayerDto) {
	botIndex := slices.IndexFunc(players, func(e dto.PlayerDto) bool {
		return e.Id == h.id
	})

	h.info = fromPlayerDto(&players[botIndex])
}

func (h *HokmBot) setTeammateInfo(players []dto.PlayerDto) {
	teammateIndex := slices.IndexFunc(players, func(e dto.PlayerDto) bool {
		return e.Id != h.id && e.Team == h.info.Team
	})
	h.teammate = fromPlayerDto(&players[teammateIndex])
}

func (h *HokmBot) handleTrumpCallerFirstCardEvent(e *event.TrumpCallerFirstCardEvent) {
	hand := map[int]model.Card{}
	for _, card := range e.Cards {
		hand[card.GetInt()] = card
	}
	h.hand = hand
}

func (h *HokmBot) handleDealCardEvent(e *event.DealCardEvent) {
	h.hand = e.Hand
	h.trump = e.Trump
}

func (h *HokmBot) handlePlayedCardEvent(e *event.PlayedCardEvent) {
	h.desk = append(h.desk, e.Card)
}

func (h *HokmBot) handleTurnWinnerEvent(e *event.TurnWinnerEvent) {
	h.desk = []*model.Card{}
}

func (h *HokmBot) Respond(schema any) {
	switch schema.(type) {
	case *response.SetTrumpResponse:
		h.SetTrump(schema.(*response.SetTrumpResponse))

	case *response.PlayCardResponse:
		h.PlayCard(schema.(*response.PlayCardResponse))
	}
}

func (h *HokmBot) SetTrump(resp *response.SetTrumpResponse) {
	resp.Suit = model.SPADE
}

func (h *HokmBot) PlayCard(resp *response.PlayCardResponse) {
	var candide model.Card
	if len(h.desk) == 0 {
		for _, card := range h.hand {
			candide = card
			break
		}
	} else {
		deskSuit := h.desk[0].Suit
		for _, card := range h.hand {
			if card.Suit == deskSuit {
				candide = card
				break
			}

			candide = card
		}
	}

	delete(h.hand, candide.GetInt())
	resp.Rank = candide.Rank
	resp.Suit = candide.Suit
}
