package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceToStringWithoutBracket(t *testing.T) {
	actualResult := SliceToStringWithoutBracket([]int{1, 3, 4})
	assert.Equal(t, "1,3,4", actualResult)

	actualResult = SliceToStringWithoutBracket([]string{"1", "3", "4"})
	assert.Equal(t, "1,3,4", actualResult)
}
