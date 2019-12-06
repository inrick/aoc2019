package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	oo := make(map[string][]string)
	for sc.Scan() {
		tt := strings.Split(sc.Text(), ")")
		if len(tt) != 2 {
			panic(tt)
		}
		x, y := tt[0], tt[1]
		oo[x] = append(oo[x], y)
	}

	fmt.Printf("a) %d\n", count(oo, "COM", 1))

	you, san := search(oo, "COM")
	b := you + san - 4 // Subtract common nodes
	fmt.Printf("b) %d\n", b)
}

func count(oo map[string][]string, root string, depth int) int {
	next := oo[root]
	n := 0
	for _, x := range next {
		n += depth + count(oo, x, depth+1)
	}
	return n
}

func search(oo map[string][]string, root string) (int, int) {
	switch root {
	case "YOU":
		return 1, 0
	case "SAN":
		return 0, 1
	}
	next := oo[root]
	var you, san int
	for _, x := range next {
		n, m := search(oo, x)
		if n > 0 && m > 0 {
			return n, m
		}
		if n > 0 {
			you = n + 1
		}
		if m > 0 {
			san = m + 1
		}
	}
	return you, san
}
