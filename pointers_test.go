package xtd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/xtd"
)

func TestFromPointer(t *testing.T) {
	var (
		t1        = "hello"
		t2 uint64 = 42
		t3        = true
		t4 *int
	)

	t.Run("string pointer", testFromPointer(&t1, true, false))
	t.Run("uint64 pointer", testFromPointer(&t2, true, false))
	t.Run("bool pointer", testFromPointer(&t3, true, false))
	t.Run("nil pointer", testFromPointer(t4, false, true))
	t.Run("pointer to nil pointer", testFromPointer(&t4, true, true))
}

func testFromPointer[T comparable](val *T, wantOk, wantZero bool) func(*testing.T) {
	return func(t *testing.T) {
		got, ok := xtd.FromPointer(val)
		if wantOk {
			assert.True(t, ok)
		} else {
			assert.False(t, ok)
		}

		if wantZero {
			assert.Zero(t, got)
		} else {
			assert.NotZero(t, got)
		}
	}
}

func TestToPointer(t *testing.T) {
	assert.NotNil(t, xtd.ToPointer("carl"))
}
