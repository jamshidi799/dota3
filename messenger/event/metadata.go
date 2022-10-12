package event

import "time"

type Metadata struct {
	Time time.Time `json:"time"`
	Type string    `json:"type"`
}

func newMetadata(t string) *Metadata {
	return &Metadata{Time: time.Now(), Type: t}
}
