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

type trumpCallerFirstCardEvent struct {
	Meta *Metadata `json:"meta"`

	Cards []model.Card `json:"cards"`
}

func NewTrumpCallerFirstCardEvent(cards []model.Card) *trumpCallerFirstCardEvent {
	return &trumpCallerFirstCardEvent{
		Meta:  newMetadata("callerFirstCard"),
		Cards: cards,
	}
}

type dealCardEvent struct {
	Meta *Metadata `json:"meta"`

	Trump model.Suit   `json:"trump"`
	Hand  []model.Card `json:"hand"`
}

func NewDealCardEvent(trump model.Suit, hand []model.Card) *dealCardEvent {
	return &dealCardEvent{
		Meta:  newMetadata("dealCard"),
		Trump: trump,
		Hand:  hand,
	}
}

type playedCardEvent struct {
	Meta *Metadata `json:"meta"`

	Card        *model.Card `json:"card"`
	PlayerIndex int         `json:"playerIndex"`
}

func NewPlayedCardEvent(card *model.Card, playerIndex int) *playedCardEvent {
	return &playedCardEvent{
		Meta:        newMetadata("playedCard"),
		Card:        card,
		PlayerIndex: playerIndex,
	}
}

type turnWinnerEvent struct {
	Meta *Metadata `json:"meta"`

	WinnerPlayerPos int `json:"winnerPlayerPos"`
}

func NewTurnWinnerEvent(winnerPlayerPos int) *turnWinnerEvent {
	return &turnWinnerEvent{
		Meta:            newMetadata("turnWinner"),
		WinnerPlayerPos: winnerPlayerPos,
	}
}

type winnerTeamEvent struct {
	Meta *Metadata `json:"meta"`

	FirstTeam  int `json:"firstTeam"`
	SecondTeam int `json:"secondTeam"`
}

func NewWinnerTeamEvent(firstTeam int, secondTeam int) *winnerTeamEvent {
	return &winnerTeamEvent{
		Meta:       newMetadata("winnerTeam"),
		FirstTeam:  firstTeam,
		SecondTeam: secondTeam,
	}
}
