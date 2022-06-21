package xtd

import (
	"context"
)

type (
	// Fn is any function which takes no inputs
	// and returns no outputs.
	Fn func()

	// NiladicFn is any function which takes no inputs
	// and returns a single output.
	NiladicFn[T any] func() T

	// NiladicErrFn is any function which takes no inputs
	// and returns a single output along with an error.
	NiladicErrFn[T any] func() (T, error)

	// UnaryFn is any function which takes a
	// single input of type T, and returns a single output.
	UnaryFn[T, U any] func(T) U

	// UnaryCtxFn is any function which takes a context.Context
	// and single input, and returns a single output.
	UnaryCtxFn[T, U any] func(context.Context, T) U

	// ErrFn is any function which takes no inputs
	// and returns an error.
	ErrFn func() error

	// UnaryErrFn is any function which takes a
	// single input, and returns a single output along with an error.
	UnaryErrFn[T, U any] func(T) (U, error)

	// UnaryCtxErrFn is any function which takes a context.Context
	// and single input, then returns a single output along with an error.
	UnaryCtxErrFn[T, U any] func(context.Context, T) (U, error)
)

// ErrFn "wraps" a Fn in a function with a ErrFn signature.
// The error returned by the wrapper function will always be nil.
func (fn Fn) ErrFn() ErrFn {
	return func() error {
		fn()
		return nil
	}
}

// ErrFn "wraps" a NiladicFn in a function with a NiladicErrFn signature.
// The error returned by the wrapper function will always be nil.
func (fn NiladicFn[T]) ErrFn() NiladicErrFn[T] {
	return func() (T, error) {
		return fn(), nil
	}
}

// ErrFn "wraps" a UnaryFn in a function with a UnaryErrFn signature.
// The error returned by the wrapper function will always be nil.
func (fn UnaryFn[T, U]) ErrFn() UnaryErrFn[T, U] {
	return func(data T) (U, error) {
		return fn(data), nil
	}
}

// ErrFn "wraps" a UnaryCtxFn in a function with a UnaryCtxErrFn signature.
// The error returned by the wrapper function will always be nil.
func (fn UnaryCtxFn[T, U]) ErrFn() UnaryCtxErrFn[T, U] {
	return func(ctx context.Context, data T) (U, error) {
		return fn(ctx, data), nil
	}
}
