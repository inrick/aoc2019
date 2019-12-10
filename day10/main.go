package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var board [][]byte
	for sc.Scan() {
		var row []byte
		for _, p := range []byte(sc.Text()) {
			switch p {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, 1)
			default:
				panic(p)
			}
		}
		board = append(board, row)
	}

	X, Y := len(board[0]), len(board)

	reach := make([][]int, Y)
	for i := range reach {
		reach[i] = make([]int, X)
	}
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if board[y][x] == 0 {
				continue
			}

			// Cast rays in all possible directions
			for ry := -Y; ry < Y; ry++ {
				for rx := -X; rx < X; rx++ {
					if d := gcd(rx, ry); (d != 1 && d != -1) || (rx == 0 && ry == 0) {
						continue
					}
					for k := 1; ; k++ {
						nx, ny := x+k*rx, y+k*ry
						if nx < 0 || ny < 0 || nx >= X || ny >= Y {
							break
						}
						if board[ny][nx] == 1 {
							reach[y][x]++
							break
						}
					}
				}
			}
		}
	}

	max, posx, posy := 0, -1, -1
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if reach[y][x] > max {
				max, posx, posy = reach[y][x], x, y
			}
		}
	}
	fmt.Printf("a) %d\n", max)

	destroyed := 0
laser:
	for {
		// Generate possible candidates for one turn then sort them by angle to get
		// the order in which the ray would have hit them. Candidates are stored
		// relative to our asteroid base.
		var candidates []struct{ x, y int }
		for ry := -Y; ry < Y; ry++ {
			for rx := -X; rx < X; rx++ {
				if d := gcd(rx, ry); (d != 1 && d != -1) || (rx == 0 && ry == 0) {
					continue
				}
				for k := 1; ; k++ {
					nx, ny := posx+k*rx, posy+k*ry
					if nx < 0 || ny < 0 || nx >= X || ny >= Y {
						break
					}
					if board[ny][nx] == 1 {
						candidates = append(candidates, struct{ x, y int }{k * rx, k * ry})
						board[ny][nx] = 0
						break
					}
				}
			}
		}

		// Sort by angle, flip arguments to Atan2 and the comparison to match what
		// our coordinate system needs.
		sort.Slice(candidates, func(i, j int) bool {
			p, q := candidates[i], candidates[j]
			return math.Atan2(float64(p.x), float64(p.y)) >
				math.Atan2(float64(q.x), float64(q.y))
		})
		for _, c := range candidates {
			destroyed++
			if destroyed == 200 {
				x, y := posx+c.x, posy+c.y
				fmt.Printf("b) %d\n", 100*x+y)
				break laser
			}
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
