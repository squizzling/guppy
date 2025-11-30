package flow

import (
	"github.com/squizzling/guppy/pkg/flow/stream"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type Published struct {
	itypes.Object

	Streams []*stream.StreamMethodPublish
}

func NewPublished() *Published {
	return &Published{
		Object: itypes.NewObject(nil),
	}
}

func (p *Published) Append(s *stream.StreamMethodPublish) {
	p.Streams = append(p.Streams, s)
}
