package main

import (
	"fmt"
)

func lpf(n int) int {
	factors := []int{}
	d := 2
	for n > 1 {
		if n%d == 0 {
			n = n / d
			factors = append(factors, d)
		}
		d += 1
		if d*d > n {
			if n > 1 {
				factors = append(factors, n)
			}
			break
		}
	}
	fmt.Println(factors)
	return factors[len(factors)-1]
}

func main() {
	fmt.Println(lpf(600851475143))
}
