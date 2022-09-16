package event

type errorEvent struct {
	Meta *Metadata

	message string
}

func NewErrorEvent(message string) *errorEvent {
	return &errorEvent{
		Meta:    newMetadata("Error"),
		message: message,
	}
}
