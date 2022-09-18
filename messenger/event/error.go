package event

type errorEvent struct {
	Meta *Metadata

	Message string
}

func NewErrorEvent(message string) *errorEvent {
	return &errorEvent{
		Meta:    newMetadata("Error"),
		Message: message,
	}
}