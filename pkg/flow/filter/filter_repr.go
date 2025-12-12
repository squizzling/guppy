package filter

import (
	"fmt"
	"strings"
)

type reprer struct{}

func mustStr(a any, err error) string {
	// repr is not allowed to fail, we'll panic if this isn't what we expect
	if err != nil {
		panic(err)
	} else {
		return a.(string)
	}
}

func repr(f Filter) string {
	return mustStr(f.Accept(reprer{}))
}

func (r reprer) VisitFilterNot(fn *FilterNot) (any, error) {
	return fmt.Sprintf("(not %s)", mustStr(fn.Right.Accept(r))), nil
}

func (r reprer) VisitFilterKeyValue(fkv *FilterKeyValue) (any, error) {
	var sb strings.Builder
	sb.WriteString("filter('")
	sb.WriteString(fkv.Key)
	sb.WriteString("'")
	for _, v := range fkv.Values {
		sb.WriteString(", '")
		sb.WriteString(v)
		sb.WriteString("'")
	}

	if fkv.MatchMissing {
		sb.WriteString(", match_missing=True")
	}
	sb.WriteString(")")
	return sb.String(), nil
}

func (r reprer) VisitFilterPartition(fp *FilterPartition) (any, error) {
	return fmt.Sprintf("partition_filter(%d, %d)", fp.Index, fp.Total), nil
}

func (r reprer) VisitFilterAnd(fa *FilterAnd) (any, error) {
	var s []string
	for _, f := range fa.Filters {
		s = append(s, mustStr(f.Accept(r)))
	}
	return fmt.Sprintf("(%s)", strings.Join(s, " and ")), nil

}

func (r reprer) VisitFilterOr(fo *FilterOr) (any, error) {
	var s []string
	for _, f := range fo.Filters {
		s = append(s, mustStr(f.Accept(r)))
	}
	return fmt.Sprintf("(%s)", strings.Join(s, " or ")), nil
}
