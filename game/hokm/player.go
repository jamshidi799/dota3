package hokm

type Player struct {
	Team          int
	position      int
	Hand          *Hand
	IsTrumpCaller bool
}

func NewPlayer(position, team int) *Player {
	return &Player{
		Team:          team,
		position:      position,
		Hand:          NewHand(),
		IsTrumpCaller: false,
	}
}
