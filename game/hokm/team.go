package hokm

type Team struct {
	Score int
}

func NewTeam() *Team {
	return &Team{
		Score: 0,
	}
}
