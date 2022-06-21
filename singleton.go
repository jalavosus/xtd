package xtd

import (
	"sync"
)

// NewSingletonFn is a type constraint allowing
// either a NiladicFn or NiladicErrFn to be passed
// to NewSingleton.
type NewSingletonFn[T any] interface {
	NiladicFn[*T] | NiladicErrFn[*T]
}

// Singleton is a container struct
// for "singleton" instances of objects.
type Singleton[T any] struct {
	val  *T
	fn   NiladicErrFn[*T]
	once *sync.Once
}

// NewSingleton constructs a Singleton with no initialized instance.
func NewSingleton[T any, Fn NewSingletonFn[T]](fn Fn) *Singleton[T] {
	var errFn NiladicErrFn[*T]

	switch f := (any)(fn).(type) {
	case NiladicFn[*T]:
		errFn = f.ErrFn()
	case NiladicErrFn[*T]:
		errFn = f
	}

	return &Singleton[T]{
		fn:   errFn,
		once: new(sync.Once),
	}
}

// Init creates a singleton instance of type T
// using the Singleton's NewSingletonFn.
// If that function returns an error, the Singleton's
// internal instance value remains nil and said error is returned.
func (s *Singleton[T]) Init() (err error) {
	s.once.Do(func() {
		var val *T

		val, err = s.fn()
		if err == nil {
			s.val = val
		}
	})

	return
}

// Val returns the Singleton's internal instance value,
// or nil if Init has not yet been called.
// Val will always return the same value after Init
// has been called.
func (s *Singleton[T]) Val() *T {
	return s.val
}

// Reset sets the Singleton's internal instance value to nil,
// and sets its internal sync.Once to a new sync.Once instance,
// effectively "resetting" the Singleton container and allowing for
// the next call to Init to have an effect.
// Do not call Reset lightly.
func (s *Singleton[T]) Reset() {
	s.once = new(sync.Once)
	s.val = nil
}
