package xtd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/xtd"
)

var (
	sliceUniqStringArgs = []string{"hello world", "hello", "hello", "world"}
	sliceUniqStringWant = []string{"hello world", "hello", "world"}

	sliceUniqUintArgs = []uint{1, 2, 3, 4, 4, 7, 1}
	sliceUniqUintWant = []uint{1, 2, 3, 4, 7}

	sliceUniqBoolArgs = []bool{true, true, false, false, true}
	sliceUniqBoolWant = []bool{true, false}
)

func TestSliceUniq(t *testing.T) {
	t.Run("string", testSliceUniq(sliceUniqStringArgs, sliceUniqStringWant))
	t.Run("uint", testSliceUniq(sliceUniqUintArgs, sliceUniqUintWant))
	t.Run("bool", testSliceUniq(sliceUniqBoolArgs, sliceUniqBoolWant))
}

func TestSliceUniqSafe(t *testing.T) {
	t.Run("string", testSliceUniqSafe(sliceUniqStringArgs, sliceUniqStringWant))
	t.Run("uint", testSliceUniqSafe(sliceUniqUintArgs, sliceUniqUintWant))
	t.Run("bool", testSliceUniqSafe(sliceUniqBoolArgs, sliceUniqBoolWant))
}

func testSliceUniq[T comparable](args, want []T) func(*testing.T) {
	return func(t *testing.T) {
		got := xtd.SliceUniq(args)
		assert.ElementsMatch(t, want, got)
	}
}

func testSliceUniqSafe[T comparable](args, want []T) func(*testing.T) {
	return func(t *testing.T) {
		got := xtd.SliceUniqSafe(args)
		assert.Equal(t, want, got)
	}
}
