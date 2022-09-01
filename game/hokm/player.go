package hokm

type Player struct {
	Id            int
	Team          Team
	position      int
	Hand          *Hand
	IsTrumpCaller bool
}

func NewPlayer(id, position, team int) *Player {
	return &Player{
		Id:            id,
		Team:          Team(team),
		position:      position,
		Hand:          NewHand(),
		IsTrumpCaller: false,
	}
}
