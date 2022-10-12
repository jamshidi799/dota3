package client

type BotClient struct {
	BaseClient
}

func (b *BotClient) GetId() int {
	return b.id
}

func (b *BotClient) write(event any) error {
	return nil
}

func (b *BotClient) read(schema any) {

}
