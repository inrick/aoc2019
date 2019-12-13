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

	sa := Init(mem, 1)
	for sa.op != 99 {
		run(&sa)
	}
	fmt.Printf("a) %d\n", sa.output)

	sb := Init(mem, 2)
	for sb.op != 99 {
		run(&sb)
	}
	fmt.Printf("b) %d\n", sb.output)
}

type state struct {
	mem    []int
	ip     int
	op     int
	rel    int
	input  int
	output int
}

func Init(mem []int, input int) state {
	mem0 := make([]int, len(mem)+(1<<10))
	copy(mem0, mem)
	return state{
		mem:    mem0,
		ip:     0,
		op:     0,
		rel:    0,
		input:  input,
		output: 0,
	}
}

func run(s *state) {
	mem := s.mem
	var modes [3]int

exec:
	for {
		mode := mem[s.ip] / 100
		s.op = mem[s.ip] % 100
		for i := range modes {
			modes[i] = mode % 10
			mode /= 10
		}
		addr := func(j int) int {
			x := mem[s.ip+1+j]
			switch modes[j] {
			case 0:
				return x
			case 2:
				return x + s.rel
			default:
				panic(mode)
			}
		}
		param := func(j int) int {
			x := mem[s.ip+1+j]
			switch modes[j] {
			case 0:
				return mem[x]
			case 1:
				return x
			case 2:
				return mem[x+s.rel]
			default:
				panic(mode)
			}
		}
		switch s.op {
		case 1:
			mem[addr(2)] = param(0) + param(1)
			s.ip += 4
		case 2:
			mem[addr(2)] = param(0) * param(1)
			s.ip += 4
		case 3:
			mem[addr(0)] = s.input
			s.ip += 2
		case 4:
			s.output = param(0)
			s.ip += 2
			return
		case 5:
			if param(0) != 0 {
				s.ip = param(1)
			} else {
				s.ip += 3
			}
		case 6:
			if param(0) == 0 {
				s.ip = param(1)
			} else {
				s.ip += 3
			}
		case 7:
			mem[addr(2)] = bool2int(param(0) < param(1))
			s.ip += 4
		case 8:
			mem[addr(2)] = bool2int(param(0) == param(1))
			s.ip += 4
		case 9:
			s.rel += param(0)
			s.ip += 2
		case 99:
			break exec
		default:
			panic(s.op)
		}
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
