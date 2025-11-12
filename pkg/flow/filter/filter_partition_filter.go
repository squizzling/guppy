package filter

import (
	"fmt"

	"guppy/pkg/interpreter"
	"guppy/pkg/interpreter/itypes"
)

type FFIPartitionFilter struct {
	itypes.Object
}

func (f FFIPartitionFilter) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "index"},
			{Name: "total"},
		},
	}, nil
}

func (f FFIPartitionFilter) resolveIndex(i itypes.Interpreter) (int, error) {
	if index, err := itypes.ArgAs[*interpreter.ObjectInt](i, "index"); err != nil {
		return 0, err
	} else {
		return index.Value, nil
	}
}

func (f FFIPartitionFilter) resolveTotal(i itypes.Interpreter) (int, error) {
	if index, err := itypes.ArgAs[*interpreter.ObjectInt](i, "total"); err != nil {
		return 0, err
	} else {
		return index.Value, nil
	}
}

func (f FFIPartitionFilter) Call(i itypes.Interpreter) (itypes.Object, error) {
	if index, err := f.resolveIndex(i); err != nil {
		return nil, err
	} else if total, err := f.resolveTotal(i); err != nil {
		return nil, err
	} else {
		return NewPartitionFilter(index, total), nil
	}
}

type partitionFilter struct {
	itypes.Object

	index int
	total int
}

func NewPartitionFilter(index int, total int) Filter {
	return &partitionFilter{
		Object: newFilterObject(),
		index:  index,
		total:  total,
	}
}

func (fpf *partitionFilter) RenderFilter() string {
	return fmt.Sprintf("partition_filter(%d, %d)", fpf.index, fpf.total)
}
