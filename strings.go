package xtd

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// ReverseString returns the passed input string, but reversed.
func ReverseString(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}

	return
}

// IntFromString wraps strconv.ParseInt, returning any of int(8/16/32/64)
// based on the specified type constraint.
func IntFromString[T constraints.Signed](s string) (res T, err error) {
	var n int64

	n, err = strconv.ParseInt(s, 10, bitSizeSigned(res))
	if err != nil {
		return
	}

	res = T(n)

	return
}

// UintFromString wraps strconv.ParseUint, returning any of uint(8/16/32/64)
// based on the specified type constraint.
func UintFromString[T constraints.Unsigned](s string) (res T, err error) {
	var n uint64

	n, err = strconv.ParseUint(s, 10, bitSizeUnsigned(res))
	if err != nil {
		return
	}

	res = T(n)

	return
}

// FloatFromString wraps strconv.ParseFloat, returning any of float(32/64)
// based on the specified type constraint.
func FloatFromString[T constraints.Float](s string) (res T, err error) {
	var f float64

	f, err = strconv.ParseFloat(s, bitSizeFloat(res))
	if err != nil {
		return
	}

	res = T(f)

	return
}

func bitSizeSigned[T constraints.Signed](n T) (bitSize int) {
	switch (any)(n).(type) {
	case int:
		bitSize = 0
	case int8:
		bitSize = 8
	case int16:
		bitSize = 16
	case int32:
		bitSize = 32
	case int64:
		bitSize = 64
	}

	return
}

func bitSizeUnsigned[T constraints.Unsigned](n T) (bitSize int) {
	switch (any)(n).(type) {
	case uint:
		bitSize = 0
	case uint8:
		bitSize = 8
	case uint16:
		bitSize = 16
	case uint32:
		bitSize = 32
	case uint64:
		bitSize = 64
	}

	return
}

func bitSizeFloat[T constraints.Float](n T) (bitSize int) {
	switch (any)(n).(type) {
	case float32:
		bitSize = 32
	case float64:
		bitSize = 64
	}

	return
}
