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

	{
		phases := []int{0, 1, 2, 3, 4}
		max := 0
		for {
			input := 0
			for _, p := range phases {
				_, input, _ = run(mem, 0, true, p, input)
			}
			if input > max {
				max = input
			}
			if !permute(phases) {
				break
			}
		}

		fmt.Printf("a) %d\n", max)
	}

	{
		phases := []int{5, 6, 7, 8, 9}
		max := 0
		for {
			var input int
			var done bool
			next := make([]cont, 5)
			for i, p := range phases {
				done, input, next[i] = run(mem, 0, true, p, input)
			}
			ndone := 0
			for i := 0; ndone < 5; i = (i + 1) % 5 {
				done, input, next[i] = next[i](input)
				if done {
					ndone++
				}
			}
			if input > max {
				max = input
			}
			if !permute(phases) {
				break
			}
		}

		fmt.Printf("b) %d\n", max)
	}
}

func permute(nn []int) bool {
	i := len(nn) - 2
	for 0 <= i && nn[i] >= nn[i+1] {
		i--
	}
	if i == -1 {
		reverse(nn)
		return false
	}
	j := len(nn) - 1
	for i <= j && nn[i] >= nn[j] {
		j--
	}
	nn[i], nn[j] = nn[j], nn[i]
	reverse(nn[i+1:])
	return true
}

func reverse(nn []int) {
	N := len(nn)
	for i := 0; i < N/2; i++ {
		nn[i], nn[N-1-i] = nn[N-1-i], nn[i]
	}
}

type cont func(int) (bool, int, cont)

func run(mem0 []int, i int, sendPhase bool, phase, myInput int) (bool, int, cont) {
	mem := make([]int, len(mem0))
	copy(mem, mem0)

	var modes [3]int
	output := myInput // Hack to get back proper value in the end

	ninput := 0
	done := false

exec:
	for {
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
			if sendPhase && ninput == 0 {
				mem[dst] = phase
			} else {
				mem[dst] = myInput
			}
			ninput++
			i += 2
		case 4:
			src := mem[i+1]
			output = value(mem, modes[0], src)
			i += 2
			return false, output, func(input int) (bool, int, cont) {
				return run(mem, i, false, 0, input)
			}
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
			done = true
			break exec
		default:
			panic(op)
		}
	}

	return done, output, func(input int) (bool, int, cont) {
		return run(mem, i, false, 0, input)
	}
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
