package main

import (
	"fmt"
)

func main() {
	const lo, hi = 359282, 820401

	a, b := 0, 0
	var nn [6]int
Search:
	for n := lo; n <= hi; n++ {
		split(&nn, n)
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

func split(nn *[6]int, n int) {
	for i := 0; i < 6; i++ {
		nn[5-i] = n % 10
		n /= 10
	}
}
