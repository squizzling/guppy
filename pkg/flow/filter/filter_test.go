package filter_test

import (
	"testing"
)

func TestFilter(t *testing.T) {
	testFromFile(t, "testdata/filter.txt")
	testFromFile(t, "testdata/filter-test.txt")
}
