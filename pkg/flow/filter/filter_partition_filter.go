package filter

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIPartitionFilter struct {
	interpreter.Object
}

func (f FFIPartitionFilter) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "index"},
			{Name: "total"},
		},
	}, nil
}

func (f FFIPartitionFilter) resolveIndex(i *interpreter.Interpreter) (int, error) {
	if index, err := interpreter.ArgAsLong(i, "index"); err != nil {
		return 0, err
	} else {
		return index, nil
	}
}

func (f FFIPartitionFilter) resolveTotal(i *interpreter.Interpreter) (int, error) {
	if index, err := interpreter.ArgAsLong(i, "total"); err != nil {
		return 0, err
	} else {
		return index, nil
	}
}

func (f FFIPartitionFilter) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if index, err := f.resolveIndex(i); err != nil {
		return nil, err
	} else if total, err := f.resolveTotal(i); err != nil {
		return nil, err
	} else {
		return NewPartitionFilter(index, total), nil
	}
}

type partitionFilter struct {
	interpreter.Object

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
