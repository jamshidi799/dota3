package hokm

type Team struct {
	Player1 *Player
	Player2 *Player
	Score   int
}

func NewTeam(pid1, pid2 string) *Team {
	return &Team{
		Player1: NewPlayer(pid1),
		Player2: NewPlayer(pid2),
		Score:   0,
	}
}
