# xtd

An e**x**tended set of things which should be in the Go s**td**lib.

## Installation

`go get github.com/jalavosus/xtd`

## Usage

```go
package main

import (
	"fmt"
	"github.com/jalavosus/xtd"
)

func main() {
	x := "hello, world"
	y := xtd.ToPointer(x)
	
	fmt.Println(x)         // "hello, world"
	fmt.Println(y)         // pointer address 
	fmt.Println(y == nil)  // false
	fmt.Println(*y == x)   // true
	
	z, ok := xtd.FromPointer(y)
	fmt.Println(ok)        // true (y is not nil)
	fmt.Println(z)         // "hello, world"
	fmt.Println(z == x)    // true
	
	var n *int
	n1, ok := xtd.FromPointer(n)
	fmt.Println(ok)        // false (n is a nil pointer)
	fmt.Println(n1)        // 0 (FromPointer returns the zero value of a type if a nil pointer is passed)
	
	uintSlice := []uint{1, 2, 3, 4, 4, 7, 1}
	uniqSlice := xtd.SliceUniq(uintSlice)
	
	fmt.Println(uintSlice) // [1 2 3 4 4 7 1]
	fmt.Println(uniqSlice) // [1 2 3 4 7] (potentially out of order)
	// Or, to ensure order is preserved:
	uniqSlice = xtd.SliceUniqSafe(uintSlice)
	fmt.Println(uniqSlice) // [1 2 3 4 7] (guaranteed to be in order respective to the input)
}
```

## TODO

- [ ] Finish writing tests
- [ ] Write GoDoc
- [ ] More functions!