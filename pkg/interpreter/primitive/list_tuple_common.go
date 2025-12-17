package primitive

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type paramSubscript struct {
	None *ObjectNone
	Int  *ObjectInt
}

func subscript(items []itypes.Object, start int) (itypes.Object, error) {
	if start < 0 {
		start = len(items) + start
	}
	if len(items) < start+1 || start < 0 {
		// TODO: Flow supports x[-1]
		return nil, fmt.Errorf("index %d out of range (0 - %d)", start, len(items)-1)
	} else {
		return items[start], nil
	}
}

func subscriptRange(items []itypes.Object, pStart paramSubscript, pEnd paramSubscript) []itypes.Object {
	var start int
	if pStart.None != nil {
		start = 0
	} else if pStart.Int.Value < 0 {
		start = len(items) + pStart.Int.Value
	} else {
		start = pStart.Int.Value
	}

	var end int
	if pEnd.None != nil {
		end = len(items)
	} else if pEnd.Int.Value < 0 {
		end = len(items) + pEnd.Int.Value
	} else {
		end = pEnd.Int.Value
	}

	// No IndexError for range
	start = clamp(0, start, len(items))
	end = clamp(0, end, len(items))

	if end < start {
		return nil
	} else {
		return items[start:end]
	}
}
