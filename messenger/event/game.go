package event

import (
	"game/model"
)

type gameStartedEvent struct {
	Meta *Metadata

	TrumpCaller int
}

func NewGameStartedEvent(trumpCaller int) *gameStartedEvent {
	return &gameStartedEvent{
		Meta:        newMetadata("GameStarted"),
		TrumpCaller: trumpCaller,
	}
}

type trumpCallerFirstCardEvent struct {
	Meta *Metadata

	Cards []model.Card
}

func NewTrumpCallerFirstCardEvent(cards []model.Card) *trumpCallerFirstCardEvent {
	return &trumpCallerFirstCardEvent{
		Meta:  newMetadata("CallerFirstCard"),
		Cards: cards,
	}
}

type dealCardEvent struct {
	Meta *Metadata

	Trump model.Suit
	Hand  []model.Card
}

func NewDealCardEvent(trump model.Suit, hand []model.Card) *dealCardEvent {
	return &dealCardEvent{
		Meta:  newMetadata("DealCard"),
		Trump: trump,
		Hand:  hand,
	}
}

type playedCardEvent struct {
	Meta *Metadata

	Card        *model.Card
	PlayerIndex int
}

func NewPlayedCardEvent(card *model.Card, playerIndex int) *playedCardEvent {
	return &playedCardEvent{
		Meta:        newMetadata("playedCard"),
		Card:        card,
		PlayerIndex: playerIndex,
	}
}

type turnWinnerEvent struct {
	Meta *Metadata

	WinnerPlayerPos int
}

func NewTurnWinnerEvent(winnerPlayerPos int) *turnWinnerEvent {
	return &turnWinnerEvent{
		Meta:            newMetadata("turnWinner"),
		WinnerPlayerPos: winnerPlayerPos,
	}
}

type winnerTeamEvent struct {
	Meta *Metadata

	FirstTeam  int
	SecondTeam int
}

func NewWinnerTeamEvent(firstTeam int, secondTeam int) *winnerTeamEvent {
	return &winnerTeamEvent{
		Meta:       newMetadata("winnerTeam"),
		FirstTeam:  firstTeam,
		SecondTeam: secondTeam,
	}
}
