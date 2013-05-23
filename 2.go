package main

import (
	"fmt"
)

func fib() func() int {
	a := 0
	b := 1

	return func() int {
		r := a + b
		a = b
		b = r
		return r
	}
}

func main() {
	r := 0
	fibber := fib()
	for {
		next := fibber()
		if next >= 4000000 {
			break
		}
		if next%2 == 0 {
			r += next
		}
	}
	fmt.Println(r)
}
