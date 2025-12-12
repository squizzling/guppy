package filter

import (
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
)

type VisitorFilter interface {
	VisitFilterNot(fn *FilterNot) (any, error)
	VisitFilterKeyValue(fkv *FilterKeyValue) (any, error)
	VisitFilterPartition(fp *FilterPartition) (any, error)
	VisitFilterAnd(fa *FilterAnd) (any, error)
	VisitFilterOr(fo *FilterOr) (any, error)
}

type Filter interface {
	itypes.Object
	Accept(vf VisitorFilter) (any, error)
}

type FilterNot struct {
	itypes.Object
	Right Filter
}

func NewFilterNot(
	Object itypes.Object,
	Right Filter,
) *FilterNot {
	return &FilterNot{
		Object: Object,
		Right:  Right,
	}
}

func (fn *FilterNot) Accept(vf VisitorFilter) (any, error) {
	return vf.VisitFilterNot(fn)
}

type FilterKeyValue struct {
	itypes.Object
	Key          string
	Values       []string
	MatchMissing bool
}

func NewFilterKeyValue(
	Object itypes.Object,
	Key string,
	Values []string,
	MatchMissing bool,
) *FilterKeyValue {
	return &FilterKeyValue{
		Object:       Object,
		Key:          Key,
		Values:       Values,
		MatchMissing: MatchMissing,
	}
}

func (fkv *FilterKeyValue) Accept(vf VisitorFilter) (any, error) {
	return vf.VisitFilterKeyValue(fkv)
}

type FilterPartition struct {
	itypes.Object
	Index int
	Total int
}

func NewFilterPartition(
	Object itypes.Object,
	Index int,
	Total int,
) *FilterPartition {
	return &FilterPartition{
		Object: Object,
		Index:  Index,
		Total:  Total,
	}
}

func (fp *FilterPartition) Accept(vf VisitorFilter) (any, error) {
	return vf.VisitFilterPartition(fp)
}

type FilterAnd struct {
	itypes.Object
	Filters []Filter
}

func NewFilterAnd(
	Object itypes.Object,
	Filters []Filter,
) *FilterAnd {
	return &FilterAnd{
		Object:  Object,
		Filters: Filters,
	}
}

func (fa *FilterAnd) Accept(vf VisitorFilter) (any, error) {
	return vf.VisitFilterAnd(fa)
}

type FilterOr struct {
	itypes.Object
	Filters []Filter
}

func NewFilterOr(
	Object itypes.Object,
	Filters []Filter,
) *FilterOr {
	return &FilterOr{
		Object:  Object,
		Filters: Filters,
	}
}

func (fo *FilterOr) Accept(vf VisitorFilter) (any, error) {
	return vf.VisitFilterOr(fo)
}
