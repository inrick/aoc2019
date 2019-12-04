package main

import (
	"fmt"
	"math"
)

func main() {
	const lo, hi = 359282, 820401

	a, b := 0, 0
Search:
	for n := lo; n <= hi; n++ {
		nn := split(n)
		va, vb := false, false
		for i := 0; i < len(nn); {
			j := i
			c := nn[j]
			for j < len(nn) && c == nn[j] {
				j++
			}
			if j < len(nn) && c > nn[j] {
				continue Search
			}
			diff := j - i
			if diff >= 2 {
				va = true
			}
			if diff == 2 {
				vb = true
			}
			i = j
		}
		if va {
			a++
		}
		if vb {
			b++
		}
	}

	fmt.Printf("a) %d\n", a)
	fmt.Printf("b) %d\n", b)
}

func split(n int) []int {
	cap := 1 + int(math.Floor(math.Log10(float64(n))))
	elems := make([]int, cap)
	for i := 0; i < cap; i++ {
		elems[cap-1-i] = n % 10
		n /= 10
	}
	return elems
}
