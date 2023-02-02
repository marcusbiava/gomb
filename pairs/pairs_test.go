package pairs

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortByValue(t *testing.T) {
	pairs := []Pair{
		{"A", 5},
		{"B", 2},
		{"C", 7},
		{"D", 1},
		{"E", 3},
	}

	expected := []Pair{
		{"D", 1},
		{"B", 2},
		{"E", 3},
		{"A", 5},
		{"C", 7},
	}

	sort.Sort(ByValue(pairs))

	if !reflect.DeepEqual(pairs, expected) {
		t.Errorf("Sorting by value failed. Expected: %v, got: %v", expected, pairs)
	}
}
