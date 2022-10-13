package bot

type Bot interface {
	Receive(msg any) error
	Respond(schema any)
}
