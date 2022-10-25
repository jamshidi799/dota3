package request

type CreateMatchRequest struct {
	BotCount int `json:"botCount"`
	WinScore int `json:"winScore"`
}
