package xtd

import (
	"context"
	"sort"
)

type sliceEntry[T any] struct {
	d   T
	idx int
}

type sliceEntries[T any] []sliceEntry[T]

func (s sliceEntries[T]) Len() int {
	return len(s)
}

func (s sliceEntries[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sliceEntries[T]) Less(i, j int) bool {
	return s[i].idx < s[j].idx
}

// SliceUniq returns a newly allocated slice containing
// only unique entries from the passed slice.
// This function does _not_ guarantee ordering of the returned slice.
func SliceUniq[T comparable](data []T) (res []T) {
	m := make(map[T]bool)

	for _, d := range data {
		if _, ok := m[d]; !ok {
			m[d] = true
		}
	}

	res = make([]T, len(m))

	idx := 0
	for d := range m {
		res[idx] = d
		idx++
	}

	return
}

// SliceUniqSafe returns a newly allocated slice containing
// only unique entries from the passed slice.
// This function does guarantee ordering of the returned slice
// respective to the passed slice, at the cost of a tad extra
// computational complexity.
func SliceUniqSafe[T comparable](data []T) (res []T) {
	m := make(map[T]int)

	var entries sliceEntries[T]

	for i, d := range data {
		if _, ok := m[d]; !ok {
			m[d] = i
			entries = append(entries, sliceEntry[T]{d, i})
		}
	}

	sort.Sort(entries)

	res = make([]T, len(entries))

	for i := range entries {
		res[i] = entries[i].d
	}

	return
}

// MapSlice calls the passed UnaryFn on all entries in the passed slice,
// returning a newly allocated slice of the same length containing that output.
func MapSlice[T, U any](data []T, fn UnaryFn[T, U]) (res []U) {
	res = make([]U, len(data))

	for i, d := range data {
		res[i] = fn(d)
	}

	return
}

// MapSliceError calls the passed UnaryErrFn on all entries in the passed slice.
// If the function returns an error after being called on any single entry,
// a nil slice and that error are returned.
// Otherwise, a newly allocated slice of the same length containing the outputs is returned.
func MapSliceError[T, U any](data []T, fn UnaryErrFn[T, U]) (res []U, err error) {
	res = make([]U, len(data))

	for i, d := range data {
		var fnRes U

		fnRes, err = fn(d)
		if err != nil {
			res = nil
			return
		}

		res[i] = fnRes
	}

	return
}

// MapSliceContext calls the passed UnaryCtxFn on all entries in the passed slice,
// returning a newly allocated slice of the same length containing that output.
func MapSliceContext[T, U any](ctx context.Context, data []T, fn UnaryCtxFn[T, U]) (res []U) {
	res = make([]U, len(data))

	for i, d := range data {
		res[i] = fn(ctx, d)
	}

	return
}

// MapSliceContextError calls the passed UnaryCtxErrFn on all entries in the passed slice.
// If the function returns an error after being called on any single entry,
// a nil slice and that error are returned.
// Otherwise, a newly allocated slice of the same length containing the outputs is returned.
func MapSliceContextError[T, U any](ctx context.Context, data []T, fn UnaryCtxErrFn[T, U]) (res []U, err error) {
	res = make([]U, len(data))

	for i, d := range data {
		var fnRes U

		fnRes, err = fn(ctx, d)
		if err != nil {
			res = nil
			return
		}

		res[i] = fnRes
	}

	return
}
