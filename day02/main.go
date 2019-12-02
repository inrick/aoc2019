package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		panic("fail")
	}
	xs := strings.Split(sc.Text(), ",")
	ops := make([]int, len(xs))
	for i, x := range xs {
		n, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		ops[i] = n
	}

	fmt.Printf("a) %d\n", run(ops, 12, 2))

	const target = 19690720

Outer:
	for x := 0; x < 99; x++ {
		for y := 0; y < 99; y++ {
			if run(ops, x, y) == target {
				fmt.Printf("b) %d\n", 100*x+y)
				break Outer
			}
		}
	}
}

func run(input []int, x, y int) int {
	N := len(input)
	ops := make([]int, N)
	copy(ops, input)
	ops[1] = x
	ops[2] = y

	for i := 0; i+3 < N; i += 4 {
		op, src0, src1, dst := ops[i], ops[i+1], ops[i+2], ops[i+3]
		switch op {
		case 1:
			ops[dst] = ops[src0] + ops[src1]
		case 2:
			ops[dst] = ops[src0] * ops[src1]
		case 99:
			break
		default:
			panic(op)
		}
	}

	return ops[0]
}
