package xtd

import (
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

func fromEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

// StringFromEnv returns a string value from the environment set
// at the given key, or the passed fallback if the key is not set.
func StringFromEnv(key string, fallback string) (val string, ok bool) {
	val, ok = fromEnv(key)
	if !ok {
		val = fallback
	}

	return
}

// IntFromEnv returns an int(8/16/32/64) value from the environment set
// at the given key, or the passed fallback if the key is not set.
func IntFromEnv[T constraints.Signed](key string, fallback T) (val T, ok bool) {
	var valStr string
	val = fallback

	valStr, ok = fromEnv(key)
	if !ok {
		return
	}

	n, err := strconv.ParseInt(valStr, 10, 64)
	if err == nil {
		val = T(n)
	}

	return
}

// UintFromEnv returns a uint(8/16/32/64) value from the environment set
// at the given key, or the passed fallback if the key is not set.
func UintFromEnv[T constraints.Unsigned](key string, fallback T) (val T, ok bool) {
	var valStr string
	val = fallback

	valStr, ok = fromEnv(key)
	if !ok {
		return
	}

	n, err := strconv.ParseUint(valStr, 10, 64)
	if err == nil {
		val = T(n)
	}

	val = T(n)

	return
}

// FloatFromEnv returns a float(32/64) value from the environment set
// at the given key, or the passed fallback if the key is not set.
func FloatFromEnv[T constraints.Float](key string, fallback T) (val T, ok bool) {
	var valStr string
	val = fallback

	valStr, ok = fromEnv(key)
	if !ok {
		return
	}

	f, err := strconv.ParseFloat(valStr, 64)
	if err == nil {
		val = T(f)
	}

	return
}

// BoolFromEnv returns a boolean value from the environment set
// at the given key, or the passed fallback if the key is not set.
func BoolFromEnv(key string, fallback bool) (val, ok bool) {
	var valStr string
	val = fallback

	valStr, ok = fromEnv(key)
	if !ok {
		return
	}

	b, err := strconv.ParseBool(valStr)
	if err == nil {
		val = b
	}

	return
}
