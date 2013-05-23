package main

import (
	"fmt"
	"sort"
)

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	for i := (len(s) - 1) / 2; i >= 0; i-- {
		if s[i:i+1] != s[len(s)-i-1:len(s)-i] {
			return false
		}
	}
	return true
}

func products(delta int) []int {
	r := []int{}
	a := 999
	b := 999 - delta
	for a >= b {
		r = append(r, a*b)
		a -= 1
		b += 1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	return sort.IntSlice(r)
}

func main() {
	for d := 0; d < 500; d++ {
		for _, v := range products(d) {
			if isPalindrome(v) {
				fmt.Println(v)
				return
			}
		}
	}
}
