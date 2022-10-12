package event

type joinPlayerEvent struct {
	Meta *Metadata `json:"meta"`

	PlayersId []int `json:"playersId"`
}

func NewJoinPlayerEvent(playerId []int) *joinPlayerEvent {
	return &joinPlayerEvent{
		Meta:      newMetadata("joinPlayer"),
		PlayersId: playerId,
	}
}
