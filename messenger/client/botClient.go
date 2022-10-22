package client

import "game/bot"

type BotClient struct {
	BaseClient
	Bot bot.Bot
}

func NewBotClient(id int, username string) *BotClient {
	return &BotClient{
		BaseClient: BaseClient{
			id:       id,
			username: username,
		},
		Bot: bot.NewHokmBot(id),
	}
}

func (b *BotClient) GetId() int {
	return b.id
}

func (b *BotClient) GetUsername() string {
	return b.username
}

func (b *BotClient) send(event any) error {
	return b.Bot.Receive(event)
}

func (b *BotClient) receive(schema any) {
	b.Bot.Respond(schema)
}
