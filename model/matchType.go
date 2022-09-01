package model

import "game/game/hokm"

func (m MatchType) GetHandler() func() error {
	if m == HOKM4 {
		return hokm.Run
	}

	return nil
}

type MatchType int

const (
	HOKM4 MatchType = iota
	HOKM3
	SHELEM
	BIDEL4
	BIDEL5
)
