package xtd

// ToPointer returns a pointer to the passed input.
func ToPointer[T any](val T) *T {
	return &val
}

// FromPointer returns the dereferenced value of the passed
// input (which is any pointer).
// If the passed input is nil, the zero value of the input's type
// is returned, along with false.
// Otherwise, the dereferenced value and true are returned.
func FromPointer[T any](val *T) (res T, ok bool) {
	if val == nil {
		return
	}

	res = *val
	ok = true

	return
}
