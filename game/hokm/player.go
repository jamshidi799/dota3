package hokm

type Player struct {
	Team          Team
	position      int
	Hand          *Hand
	IsTrumpCaller bool
}

func NewPlayer(position, team int) *Player {
	return &Player{
		Team:          Team(team),
		position:      position,
		Hand:          NewHand(),
		IsTrumpCaller: false,
	}
}
