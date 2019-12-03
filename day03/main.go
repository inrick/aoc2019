package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type v2 struct {
	x, y int
}

type move struct {
	step v2
	magn int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var paths [][]move
	for sc.Scan() {
		xs := strings.Split(sc.Text(), ",")
		var path []move
		for _, x := range xs {
			var step v2
			switch x[0] {
			case 'L':
				step = v2{-1, 0}
			case 'R':
				step = v2{1, 0}
			case 'U':
				step = v2{0, 1}
			case 'D':
				step = v2{0, -1}
			default:
				panic(nil)
			}
			magn, err := strconv.Atoi(x[1:])
			if err != nil {
				panic(err)
			}
			path = append(path, move{step, magn})
		}
		paths = append(paths, path)
	}

	var crossings []v2
	visited := make(map[v2]byte)
	steps := make(map[v2]int)
	for i, path := range paths {
		n := byte(i) + 1 // Path number
		curr := v2{0, 0}
		nsteps := 0
		visited[curr] = n
		for _, m := range path {
			// Walk
			for j := 0; j < m.magn; j++ {
				nsteps++
				curr = add(curr, m.step)
				if mark := visited[curr]; mark != n {
					if mark != 0 {
						crossings = append(crossings, curr)
					}
					steps[curr] += nsteps
				}
				visited[curr] = n
			}
		}
	}

	minDist, minSteps := math.MaxInt64, math.MaxInt64
	for _, u := range crossings {
		if d := dist(u, v2{0, 0}); d < minDist {
			minDist = d
		}
		if s := steps[u]; s < minSteps {
			minSteps = s
		}
	}

	fmt.Printf("a) %d\n", minDist)
	fmt.Printf("b) %d\n", minSteps)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(u, v v2) int {
	return abs(u.x-v.x) + abs(u.y-v.y)
}

func add(u, v v2) v2 {
	return v2{u.x + v.x, u.y + v.y}
}
