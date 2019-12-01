package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var lines []int
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, n)
	}

	a, b := 0, 0
	for _, n := range lines {
		m := n/3 - 2
		a += m
		for m > 0 {
			b += m
			m = m/3 - 2
		}
	}

	fmt.Printf("a) %d\n", a)
	fmt.Printf("b) %d\n", b)
}
