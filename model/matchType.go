package model

type MatchType struct {
	PlayerCount int
	Type GameType
}

type GameType int

const (
	HOKM GameType = iota
	SHELEM
	BIDEL
)
