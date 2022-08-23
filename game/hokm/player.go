package hokm

type Player struct {
	ID            string
	position      int
	Hand          *Hand
	IsTrumpCaller bool
}

func NewPlayer(id string) *Player {
	return &Player{
		ID:            id,
		Hand:          NewHand(),
		IsTrumpCaller: false,
	}
}
