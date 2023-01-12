package slices

import (
	"fmt"
	"strings"
)

func SliceToStringWithoutBracket[T any](slice []T) string {
	temp := fmt.Sprintf("%v", slice)
	split := strings.Split(temp[1:len(temp)-1], " ")
	join := strings.Join(split, ",")
	return fmt.Sprintf("%v", join)
}
