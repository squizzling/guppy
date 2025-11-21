package filter

import (
	"fmt"

	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type ffiPartitionFilter struct {
	Index *primitive.ObjectInt `ffi:"index"`
	Total *primitive.ObjectInt `ffi:"total"`
}

func NewFFIPartitionFilter() itypes.FlowCall {
	return ffi.NewFFI(ffiPartitionFilter{})
}

func (f ffiPartitionFilter) Call(i itypes.Interpreter) (itypes.Object, error) {
	return NewPartitionFilter(f.Index.Value, f.Total.Value), nil
}

type partitionFilter struct {
	itypes.Object

	index int
	total int
}

func NewPartitionFilter(index int, total int) Filter {
	return &partitionFilter{
		Object: prototypeFilter,
		index:  index,
		total:  total,
	}
}

func (pf *partitionFilter) Repr() string {
	return fmt.Sprintf("partition_filter(%d, %d)", pf.index, pf.total)
}
