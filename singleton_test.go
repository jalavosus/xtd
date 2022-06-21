package xtd_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/xtd"
)

const (
	testSingletonVal uint64 = 42
)

var (
	testNewSingletonFn xtd.NiladicFn[*uint64] = func() *uint64 {
		var val = testSingletonVal
		return &val
	}

	testNewSingletonErrFn xtd.NiladicErrFn[*uint64] = func() (*uint64, error) {
		var val = testSingletonVal
		return &val, fmt.Errorf("an error")
	}
)

func TestSingleton_Init(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		singleton := xtd.NewSingleton[uint64](testNewSingletonFn)
		assert.NoError(t, singleton.Init())
	})

	t.Run("with error", func(t *testing.T) {
		singleton := xtd.NewSingleton[uint64](testNewSingletonErrFn)
		assert.Error(t, singleton.Init())
		assert.NoError(t, singleton.Init())
	})
}

func TestSingleton_Reset(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		singleton := xtd.NewSingleton[uint64](testNewSingletonFn)

		assert.Nil(t, singleton.Val())

		assert.NoError(t, singleton.Init())
		assert.NotNil(t, singleton.Val())
		assert.Equal(t, testSingletonVal, *(singleton.Val()))

		singleton.Reset()
		assert.Nil(t, singleton.Val())

		assert.NoError(t, singleton.Init())
		assert.NotNil(t, singleton.Val())
		assert.Equal(t, testSingletonVal, *(singleton.Val()))
	})
}

func TestSingleton_Val(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		singleton := xtd.NewSingleton[uint64](testNewSingletonFn)

		assert.Nil(t, singleton.Val())

		assert.NoError(t, singleton.Init())

		assert.NotNil(t, singleton.Val())
		assert.Equal(t, testSingletonVal, *(singleton.Val()))
	})
}
