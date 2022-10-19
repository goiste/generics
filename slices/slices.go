package slices

import (
	"fmt"
	"math"
	"sort"
)

// Number
// represents numeric types
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Copy
// returns a copy of the slice
func Copy[T any](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	return result
}

// RemoveValue
// returns a slice with the element(s) removed by value
func RemoveValue[T comparable](s []T, value T) []T {
	result := make([]T, 0, len(s))
	for i := range s {
		if s[i] != value {
			result = append(result, s[i])
		}
	}
	return result[:len(result):len(result)]
}

// RemoveIdx
// returns a slice with the element removed by index
func RemoveIdx[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	result := Copy(s)
	return append(result[:index:len(result)-1], result[index+1:]...)
}

// HasValue
// returns true if the slice contains the value or false if not
func HasValue[T comparable](s []T, value T) bool {
	return IndexOf(s, value) != -1
}

// IndexOf
// returns index of given value or -1 if the value is not found
func IndexOf[T comparable](s []T, value T) int {
	for i := range s {
		if s[i] == value {
			return i
		}
	}
	return -1
}

// Diff
// returns a slice that contains elements that are not represented in any of the other slices
func Diff[T comparable](s []T, others ...[]T) []T {
	if len(others) == 0 {
		return s
	}

	valuesMap := make(map[T]struct{}, len(s))
	for i := range s {
		valuesMap[s[i]] = struct{}{}
	}

	result := Copy(s)
	for i := range others {
		for j := range others[i] {
			if _, exists := valuesMap[others[i][j]]; exists {
				result = RemoveValue(result, others[i][j])
			}
		}
	}

	return result
}

// Intersect
// returns a slice containing elements represented in each of the others slices
func Intersect[T comparable](s []T, others ...[]T) []T {
	if len(others) == 0 {
		return []T{}
	}

	sort.Slice(others, func(i, j int) bool {
		return len(others[i]) < len(others[j])
	})

	if len(others[0]) == 0 {
		return []T{}
	}

	result := Copy(s)
	for i := range others {
		result = Filter(result, func(t T) bool { return HasValue(others[i], t) })
		if len(result) == 0 {
			return result
		}
	}

	return result
}

// SafeSlice
// returns slice of slice via copy
func SafeSlice[T any](s []T, from, to int) []T {
	if to > len(s) {
		to = len(s)
	}

	return Copy(s[from:to])
}

// Split
// splits a given slice into parts by `partSize` elements
func Split[T any](s []T, partSize int) [][]T {
	if partSize <= 0 || len(s) == 0 {
		return [][]T{}
	}

	if len(s) <= partSize {
		return [][]T{SafeSlice(s, 0, len(s))}
	}

	length := int(math.Ceil(float64(len(s)) / float64(partSize)))

	result := make([][]T, length)

	for i := 0; i < length; i++ {
		result[i] = SafeSlice(s, i*partSize, (i+1)*partSize)
	}

	return result
}

// Merge
// returns a slice containing all unique elements of all slices
func Merge[T comparable](s []T, others ...[]T) []T {
	result := Copy(s)
	for i := range others {
		result = append(result, Diff(others[i], result)...)
	}
	return result
}

// Unique
// returns a slice containing only unique elements of the original slice
func Unique[T comparable](s []T) []T {
	result := make([]T, 0, len(s))
	tmp := map[T]struct{}{}
	for i := range s {
		if _, exists := tmp[s[i]]; exists {
			continue
		}
		result = append(result, s[i])
		tmp[s[i]] = struct{}{}
	}
	return result[:len(result):len(result)]
}

// Fill
// returns a slice of length `len` filled with the value of `t`
func Fill[T any](value T, len int) []T {
	if len <= 0 {
		return []T{}
	}

	result := make([]T, len)
	for i := 0; i < len; i++ {
		result[i] = value
	}

	return result
}

// Filter
// returns a slice containing elements for which the func returns true
func Filter[T any](s []T, f func(value T) bool) []T {
	if f == nil {
		return s
	}

	result := Copy(s)

	for i := 0; i < len(result); {
		if !f(result[i]) {
			result = RemoveIdx(result, i)
			continue
		}
		i++
	}

	return result
}

// Map
// returns a slice filled with the values returned by the func for each slice element
func Map[T any](s []T, f func(value T) T) []T {
	if f == nil {
		return s
	}

	result := make([]T, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

// Convert
// converts all elements from numeric type T to numeric type V
func Convert[T, V Number](s []T) []V {
	result := make([]V, len(s))
	for i := range s {
		result[i] = V(s[i])
	}
	return result
}

// Min
// returns the minimum element (numeric types only)
func Min[T Number](s []T) T {
	if len(s) == 0 {
		return *new(T)
	}
	result := s[0]
	for i := range s {
		if s[i] < result {
			result = s[i]
		}
	}
	return result
}

// Max
// returns the maximum element (numeric types only)
func Max[T Number](s []T) T {
	if len(s) == 0 {
		return *new(T)
	}
	result := s[0]
	for i := range s {
		if s[i] > result {
			result = s[i]
		}
	}
	return result
}

// Sum
// returns the sum of all element (numeric types only)
func Sum[T Number](s []T) T {
	result := *new(T)
	for i := range s {
		result += s[i]
	}
	return result
}

// Reverse
// returns reversed slice
func Reverse[T any](s []T) []T {
	result := Copy(s)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// Format
// returns a slice containing source elements formatted by format string
func Format[T any](s []T, format string) []string {
	result := make([]string, len(s))
	for i := range s {
		result[i] = fmt.Sprintf(format, s[i])
	}
	return result
}

// Stringify
// returns a slice containing string representations of the source elements
func Stringify[T fmt.Stringer](s []T) []string {
	result := make([]string, len(s))
	for i := range s {
		result[i] = s[i].String()
	}
	return result
}

// Range
// returns a slice of the numeric elements from `start` (include) to `stop` (exclude) with given `step`
func Range[T Number](start, stop, step T) []T {
	if step == 0 || start == stop || (start > stop && step > 0) {
		return []T{}
	}

	result := make([]T, 0, int(math.Ceil(float64(stop-start)/float64(step))))

	end := func(i T) bool { return i < stop }
	if start > stop {
		end = func(i T) bool { return i > stop }
	}

	for i := start; end(i); i += step {
		result = append(result, i)
	}

	return result
}

// SequenceGenerator
// returns a function that returns a value incremented/decremented by `step` on each call with initial value of `start`
func SequenceGenerator[T Number](start, step T) func() T {
	start -= step
	return func() T {
		result := start + step
		start += step
		return result
	}
}
