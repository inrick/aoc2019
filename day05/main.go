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
	xx := strings.Split(sc.Text(), ",")
	mem := make([]int, len(xx))
	for i, x := range xx {
		n, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		mem[i] = n
	}

	fmt.Printf("a) %d\n", run(mem, 1))
	fmt.Printf("b) %d\n", run(mem, 5))
}

func run(input []int, myInput int) int {
	mem := make([]int, len(input))
	copy(mem, input)

	var modes [3]int
	var output int

exec:
	for i := 0; ; {
		mode, op := mem[i]/100, mem[i]%100
		for i := range modes {
			modes[i] = mode % 10
			mode /= 10
		}
		switch op {
		case 1:
			src0, src1, dst := mem[i+1], mem[i+2], mem[i+3]
			mem[dst] = value(mem, modes[0], src0) + value(mem, modes[1], src1)
			i += 4
		case 2:
			src0, src1, dst := mem[i+1], mem[i+2], mem[i+3]
			mem[dst] = value(mem, modes[0], src0) * value(mem, modes[1], src1)
			i += 4
		case 3:
			dst := mem[i+1]
			mem[dst] = myInput
			i += 2
		case 4:
			src := mem[i+1]
			output = value(mem, modes[0], src)
			i += 2
		case 5:
			src, jmp := mem[i+1], mem[i+2]
			if value(mem, modes[0], src) != 0 {
				i = value(mem, modes[1], jmp)
			} else {
				i += 3
			}
		case 6:
			src, jmp := mem[i+1], mem[i+2]
			if value(mem, modes[0], src) == 0 {
				i = value(mem, modes[1], jmp)
			} else {
				i += 3
			}
		case 7:
			src0, src1, dst := mem[i+1], mem[i+2], mem[i+3]
			mem[dst] = bool2int(value(mem, modes[0], src0) < value(mem, modes[1], src1))
			i += 4
		case 8:
			src0, src1, dst := mem[i+1], mem[i+2], mem[i+3]
			mem[dst] = bool2int(value(mem, modes[0], src0) == value(mem, modes[1], src1))
			i += 4
		case 99:
			break exec
		default:
			panic(op)
		}
	}

	return output
}

// Need a form that the go compiler optimizes. See issue #6011.
func bool2int(b bool) int {
	var i int
	if b {
		i = 1
	}
	return i
}

func value(mem []int, mode, param int) int {
	switch mode {
	case 0:
		return mem[param]
	case 1:
		return param
	default:
		panic(nil)
	}
}
