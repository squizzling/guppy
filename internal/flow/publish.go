package flow

import (
	"guppy/internal/flow/stream"
	"guppy/internal/interpreter"
)

type Published struct {
	interpreter.Object

	Streams []*stream.StreamPublish
}

func NewPublished() *Published {
	return &Published{
		Object: interpreter.NewObject(nil),
	}
}

func (p *Published) Append(s *stream.StreamPublish) {
	p.Streams = append(p.Streams, s)
}
