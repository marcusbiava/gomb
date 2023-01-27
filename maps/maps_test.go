package maps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	keys := Keys(map[string]int{
		"k1": 1,
		"k2": 2,
		"k3": 3,
	})
	sort.Strings(keys)
	assert.Equal(t, "[k1 k2 k3]", fmt.Sprintf("%v", keys))
}

func TestValues(t *testing.T) {
	values := Values(map[string]int{
		"k1": 1,
		"k2": 2,
		"k3": 3,
	})
	sort.Ints(values)
	assert.Equal(t, "[1 2 3]", fmt.Sprintf("%v", values))
}
