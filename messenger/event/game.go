package event

import (
	"game/model"
	"time"
)

type gameStartedEvent struct {
	Meta Metadata

	TrumpCaller int
}

func NewGameStartedEvent(trumpCaller int) *gameStartedEvent {
	return &gameStartedEvent{
		Meta: Metadata{
			Time: time.Now(),
			Type: "GameStarted",
		},
		TrumpCaller: trumpCaller,
	}
}

type trumpCallerFirstCardEvent struct {
	Meta Metadata

	Cards []model.Card
}

func NewTrumpCallerFirstCardEvent(cards []model.Card) *trumpCallerFirstCardEvent {
	return &trumpCallerFirstCardEvent{
		Meta: Metadata{
			Time: time.Now(),
			Type: "CallerFirstCard",
		},
		Cards: cards,
	}
}
