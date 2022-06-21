package xtd_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/xtd"
)

type FromEnvTestCase[T comparable] struct {
	name   string
	args   FromEnvTestArgs[T]
	want   T
	wantOk bool
	setEnv bool
}

type FromEnvTestArgs[T comparable] struct {
	key      string
	val      string
	fallback T
}

func setEnv[T any](t *testing.T, key string, val T) {
	t.Helper()
	valStr := fmt.Sprintf("%v", val)
	t.Setenv(key, valStr)
}

func TestBoolFromEnv(t *testing.T) {
	type (
		testCase FromEnvTestCase[bool]
		testArgs = FromEnvTestArgs[bool]
	)

	tests := []testCase{
		{
			name:   "true",
			args:   testArgs{"puppies", "true", true},
			want:   true,
			wantOk: true,
			setEnv: true,
		},
		{
			name:   "fallback true",
			args:   testArgs{"puppies", "yes", true},
			setEnv: false,
			want:   true,
			wantOk: false,
		},
		{
			name:   "fallback true",
			args:   testArgs{"puppies", "yes", false},
			setEnv: true,
			want:   false,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				setEnv(t, tt.args.key, tt.args.val)
			}

			gotVal, gotOk := xtd.BoolFromEnv(tt.args.key, tt.args.fallback)
			assert.Equal(t, tt.want, gotVal, "expected gotVal to equal %v, got %v", tt.args.val, tt.args.fallback)
			assert.Equal(t, tt.wantOk, gotOk, "expected gotOk to equal %v, got %v", tt.wantOk, tt.args.fallback)
		})
	}
}

func TestFloatFromEnv(t *testing.T) {
	type (
		testCase FromEnvTestCase[float64]
		testArgs = FromEnvTestArgs[float64]
	)

	tests := []testCase{
		{
			name:   "true",
			args:   testArgs{"puppies", "42.33", 0},
			want:   42.33,
			wantOk: true,
			setEnv: true,
		},
		{
			name:   "fallback",
			args:   testArgs{"puppies", "42.323", 0},
			setEnv: false,
			want:   0,
			wantOk: false,
		},
		{
			name:   "fallback (setenv, parse error)",
			args:   testArgs{"puppies", "carl", 0},
			setEnv: true,
			want:   0,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				setEnv(t, tt.args.key, tt.args.val)
			}

			gotVal, gotOk := xtd.FloatFromEnv(tt.args.key, tt.args.fallback)

			assert.Equal(t, tt.want, gotVal, "expected gotVal to equal %v, got %v", tt.args.val, tt.args.fallback)
			assert.Equal(t, tt.wantOk, gotOk, "expected gotOk to equal %v, got %v", tt.wantOk, tt.args.fallback)
		})
	}
}

func TestIntFromEnv(t *testing.T) {
	type (
		testCase FromEnvTestCase[int]
		testArgs = FromEnvTestArgs[int]
	)

	tests := []testCase{
		{
			name:   "true",
			args:   testArgs{"puppies", "42", 0},
			want:   42,
			wantOk: true,
			setEnv: true,
		},
		{
			name:   "fallback",
			args:   testArgs{"puppies", "42", 0},
			setEnv: false,
			want:   0,
			wantOk: false,
		},
		{
			name:   "fallback (setenv, parse error)",
			args:   testArgs{"puppies", "42.323", 0},
			setEnv: true,
			want:   0,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				setEnv(t, tt.args.key, tt.args.val)
			}

			gotVal, gotOk := xtd.IntFromEnv(tt.args.key, tt.args.fallback)

			assert.Equal(t, tt.want, gotVal, "expected gotVal to equal %v, got %v", tt.args.val, tt.args.fallback)
			assert.Equal(t, tt.wantOk, gotOk, "expected gotOk to equal %v, got %v", tt.wantOk, tt.args.fallback)
		})
	}
}

func TestStringFromEnv(t *testing.T) {
	type (
		testCase FromEnvTestCase[string]
		testArgs = FromEnvTestArgs[string]
	)

	tests := []testCase{
		{
			name:   "true",
			args:   testArgs{"puppies", "42", ""},
			want:   "42",
			wantOk: true,
			setEnv: true,
		},
		{
			name:   "fallback",
			args:   testArgs{"puppies", "42", ""},
			setEnv: false,
			want:   "",
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				setEnv(t, tt.args.key, tt.args.val)
			}

			gotVal, gotOk := xtd.StringFromEnv(tt.args.key, tt.args.fallback)

			assert.Equal(t, tt.want, gotVal, "expected gotVal to equal %v, got %v", tt.args.val, tt.args.fallback)
			assert.Equal(t, tt.wantOk, gotOk, "expected gotOk to equal %v, got %v", tt.wantOk, tt.args.fallback)
		})
	}
}

func TestUintFromEnv(t *testing.T) {
	type (
		testCase FromEnvTestCase[uint]
		testArgs = FromEnvTestArgs[uint]
	)

	tests := []testCase{
		{
			name:   "true",
			args:   testArgs{"puppies", "42", 0},
			want:   42,
			wantOk: true,
			setEnv: true,
		},
		{
			name:   "fallback",
			args:   testArgs{"puppies", "42", 0},
			setEnv: false,
			want:   0,
			wantOk: false,
		},
		{
			name:   "fallback (setenv, parse error)",
			args:   testArgs{"puppies", "42.323", 0},
			setEnv: true,
			want:   0,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				setEnv(t, tt.args.key, tt.args.val)
			}

			gotVal, gotOk := xtd.UintFromEnv(tt.args.key, tt.args.fallback)

			assert.Equal(t, tt.want, gotVal, "expected gotVal to equal %v, got %v", tt.args.val, tt.args.fallback)
			assert.Equal(t, tt.wantOk, gotOk, "expected gotOk to equal %v, got %v", tt.wantOk, tt.args.fallback)
		})
	}
}
