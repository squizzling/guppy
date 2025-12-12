package filter

import (
	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
)

type ffiPartitionFilter struct {
	Index *primitive.ObjectInt `ffi:"index"`
	Total *primitive.ObjectInt `ffi:"total"`
}

func NewFFIPartitionFilter() itypes.FlowCall {
	return ffi.NewFFI(ffiPartitionFilter{})
}

func (f ffiPartitionFilter) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewFilterPartition(prototypeFilter, f.Index.Value, f.Total.Value), nil
}

func (fp *FilterPartition) Repr() string {
	return repr(fp)
}
