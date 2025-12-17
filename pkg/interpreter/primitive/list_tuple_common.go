package primitive

import (
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

func subscript(items []itypes.Object, start int) (itypes.Object, error) {
	if len(items) < start+1 || start < 0 {
		// TODO: Flow supports x[-1]
		return nil, fmt.Errorf("index %d out of range (0 - %d)", start, len(items)-1)
	} else {
		return items[start], nil
	}
}

func subscriptRange(items []itypes.Object, pStart *int, pEnd *int) []itypes.Object {
	var start int
	if pStart == nil {
		start = 0
	} else if *pStart < 0 {
		start = len(items) + *pStart
	} else {
		start = *pStart
	}

	var end int
	if pEnd == nil {
		end = len(items)
	} else if *pEnd < 0 {
		end = len(items) + *pEnd
	} else {
		end = *pEnd
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
