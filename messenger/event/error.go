package event

type errorEvent struct {
	Meta *Metadata `json:"meta"`

	Message string `json:"message"`
}

func NewErrorEvent(message string) *errorEvent {
	return &errorEvent{
		Meta:    newMetadata("error"),
		Message: message,
	}
}
