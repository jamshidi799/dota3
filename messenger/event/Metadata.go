package event

import "time"

type Metadata struct {
	Time time.Time
	Type string
}

func newMetadata(t string) *Metadata {
	return &Metadata{Time: time.Now(), Type: t}
}
