package slices

import (
	"fmt"
	"strings"
)

func SliceToStringWithoutBracket[T any](slice []T) string {
	strSlice := make([]string, len(slice))
	for i, value := range slice {
		strSlice[i] = fmt.Sprintf("%v", value)
	}
	return strings.Join(strSlice, ",")
}

func Map[T any, R any](slice []T, mapper func(value T, index int, slice []T) R) (mapped []R) {
	for i, el := range slice {
		mapped = append(mapped, mapper(el, i, slice))
	}
	return mapped
}

func Chunk[T any](input []T, size int) [][]T {
	var chunks [][]T

	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

func Filter[T any](slice []T, predicate func(value T, index int, slice []T) bool) (filtered []T) {
	for i, el := range slice {
		if ok := predicate(el, i, slice); ok {
			filtered = append(filtered, el)
		}
	}
	return filtered
}
