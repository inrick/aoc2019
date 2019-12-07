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
				mem0 := make([]int, len(mem))
				copy(mem0, mem)
				input = run(&state{mem0, 0, true, p, false}, input)
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
			input := 0
			ss := make([]state, 5)
			for i, p := range phases {
				mem0 := make([]int, len(mem))
				copy(mem0, mem)
				ss[i] = state{mem0, 0, true, p, false}
				input = run(&ss[i], input)
			}
			for i := 0; !ss[4].done; i = (i + 1) % 5 {
				input = run(&ss[i], input)
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

type state struct {
	mem      []int
	ip       int
	setPhase bool
	phase    int
	done     bool
}

func run(s *state, input int) (output int) {
	mem := s.mem
	var modes [3]int

exec:
	for i := s.ip; ; {
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
			if s.setPhase {
				mem[dst] = s.phase
				s.setPhase = false
			} else {
				mem[dst] = input
			}
			i += 2
		case 4:
			src := mem[i+1]
			output = value(mem, modes[0], src)
			i += 2
			s.ip = i
			return output
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
			s.done = true
			s.ip = i
			break exec
		default:
			panic(op)
		}
	}

	return input
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
