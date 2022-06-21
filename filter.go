package xtd

import (
	"sort"
)

// FilterFn is any function which takes
// a single input and returns a boolean.
type FilterFn[T any] func(T) bool

// FilterSlice returns a newly allocated slice containing all
// entries from the passed slice for which the passed FilterFn returns true.
// This function does not gauarantee that the ordering of entries
// in the resulting slice is in any way similar to the input slice.
func FilterSlice[T any](data []T, fn FilterFn[T]) (res []T) {
	res = make([]T, 0, len(data))

	for _, d := range data {
		if fn(d) {
			res = append(res, d)
		}
	}

	return
}

// FilterSliceSafe returns a newly allocated slice containing all
// entries from the passed slice for which the passed FilterFn returns true.
// This function does gauarantee that the ordering of entries
// in the resulting slice is respective of the ordering of the input slice,
// at the cost of a tad bit extra computational complexity.
func FilterSliceSafe[T any](data []T, fn FilterFn[T]) (res []T) {
	var entries sliceEntries[T]

	for i, d := range data {
		if fn(d) {
			entries = append(entries, sliceEntry[T]{d, i})
		}
	}

	res = make([]T, len(entries))

	sort.Sort(entries)

	for i := range entries {
		res[i] = entries[i].d
	}

	return
}
