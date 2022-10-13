package event

import (
	"game/messenger/dto"
	"game/model"
)

type GameStartedEvent struct {
	Meta *Metadata `json:"meta"`

	Players []dto.PlayerDto `json:"players"`
}

func NewGameStartedEvent(players []dto.PlayerDto) *GameStartedEvent {
	return &GameStartedEvent{
		Meta:    newMetadata("gameStarted"),
		Players: players,
	}
}

type TrumpCallerFirstCardEvent struct {
	Meta *Metadata `json:"meta"`

	Cards []model.Card `json:"cards"`
}

func NewTrumpCallerFirstCardEvent(cards []model.Card) *TrumpCallerFirstCardEvent {
	return &TrumpCallerFirstCardEvent{
		Meta:  newMetadata("callerFirstCard"),
		Cards: cards,
	}
}

type DealCardEvent struct {
	Meta *Metadata `json:"meta"`

	Trump model.Suit         `json:"trump"`
	Hand  map[int]model.Card `json:"hand"`
}

func NewDealCardEvent(trump model.Suit, hand map[int]model.Card) *DealCardEvent {
	return &DealCardEvent{
		Meta:  newMetadata("dealCard"),
		Trump: trump,
		Hand:  hand,
	}
}

type PlayedCardEvent struct {
	Meta *Metadata `json:"meta"`

	Card        *model.Card `json:"card"`
	PlayerIndex int         `json:"playerIndex"`
}

func NewPlayedCardEvent(card *model.Card, playerIndex int) *PlayedCardEvent {
	return &PlayedCardEvent{
		Meta:        newMetadata("playedCard"),
		Card:        card,
		PlayerIndex: playerIndex,
	}
}

type TurnWinnerEvent struct {
	Meta *Metadata `json:"meta"`

	WinnerPlayerPos int `json:"winnerPlayerPos"`
}

func NewTurnWinnerEvent(winnerPlayerPos int) *TurnWinnerEvent {
	return &TurnWinnerEvent{
		Meta:            newMetadata("turnWinner"),
		WinnerPlayerPos: winnerPlayerPos,
	}
}

type WinnerTeamEvent struct {
	Meta *Metadata `json:"meta"`

	FirstTeam  int `json:"firstTeam"`
	SecondTeam int `json:"secondTeam"`
}

func NewWinnerTeamEvent(firstTeam int, secondTeam int) *WinnerTeamEvent {
	return &WinnerTeamEvent{
		Meta:       newMetadata("winnerTeam"),
		FirstTeam:  firstTeam,
		SecondTeam: secondTeam,
	}
}
