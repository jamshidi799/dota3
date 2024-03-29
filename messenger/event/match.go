package event

type joinPlayerEvent struct {
	Meta *Metadata `json:"meta"`

	Players []string `json:"players"`
}

func NewJoinPlayerEvent(playersUsername []string) *joinPlayerEvent {
	return &joinPlayerEvent{
		Meta:    newMetadata("joinPlayer"),
		Players: playersUsername,
	}
}

type EndMatchEvent struct {
	Meta *Metadata `json:"meta"`
}

func NewEndMatchEvent() *EndMatchEvent {
	return &EndMatchEvent{
		Meta: newMetadata("endMatch"),
	}
}
