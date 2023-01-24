package slices

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSliceToStringWithoutBracket(t *testing.T) {
	actualResult := SliceToStringWithoutBracket([]int{1, 3, 4})
	assert.Equal(t, "1,3,4", actualResult)

	actualResult = SliceToStringWithoutBracket([]string{"1", "3", "4"})
	assert.Equal(t, "1,3,4", actualResult)
}

func TestMap(t *testing.T) {
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedResult := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	actualResult := Map(numerals, func(value int, _ int, _ []int) string {
		return strconv.Itoa(value)
	})
	assert.Equal(t, expectedResult, actualResult)
}
