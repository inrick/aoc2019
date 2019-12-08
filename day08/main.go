package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		panic("fail")
	}
	input := []byte(sc.Text())
	img := make([]byte, len(input))
	for i, b := range input {
		img[i] = b - '0'
	}

	const N = 25 * 6
	nlayers := len(img) / N
	min := [3]int{math.MaxInt64, 0, 0}
	for i := 0; i < nlayers; i++ {
		var count [3]int
		for j := 0; j < N; j++ {
			count[img[N*i+j]]++
		}
		if count[0] < min[0] {
			min = count
		}
	}

	fmt.Printf("a) %d\n", min[1]*min[2])

	render := make([]byte, N)
	for i := 0; i < N; i++ {
		render[i] = 2
	}
	for i := 0; i < nlayers; i++ {
		for j := 0; j < N; j++ {
			if render[j] == 2 {
				render[j] = img[N*i+j]
			}
		}
	}

	fmt.Println("b)")
	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			fmt.Printf("%c", " #"[render[25*y+x]])
		}
		fmt.Println()
	}
}
