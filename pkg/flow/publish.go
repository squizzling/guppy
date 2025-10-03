package flow

import (
	"guppy/pkg/flow/stream"
	"guppy/pkg/interpreter"
)

type Published struct {
	interpreter.Object

	Streams []*stream.StreamMethodPublish
}

func NewPublished() *Published {
	return &Published{
		Object: interpreter.NewObject(nil),
	}
}

func (p *Published) Append(s *stream.StreamMethodPublish) {
	p.Streams = append(p.Streams, s)
}
