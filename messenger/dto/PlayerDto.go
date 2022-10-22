package dto

type PlayerDto struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Team          int    `json:"team"`
	Position      int    `json:"position"`
	IsTrumpCaller bool   `json:"isTrumpCaller"`
}
