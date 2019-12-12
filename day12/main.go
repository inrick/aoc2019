package main

import (
	"bufio"
	"fmt"
	"os"
)

type v3 struct {
	x, y, z int
}

type moon struct {
	pos v3
	vel v3
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var moons []moon
	for sc.Scan() {
		var u v3
		fmt.Sscanf(sc.Text(), "<x=%d, y=%d, z=%d>", &u.x, &u.y, &u.z)
		moons = append(moons, moon{pos: u})
	}

	prev := make([][3][4]int, 0)

	for t := 0; t < 1000000; t++ {
		if t == 1000 {
			total := 0
			for _, m := range moons {
				total += absSum(m.pos) * absSum(m.vel)
			}
			fmt.Printf("a) %d\n", total)
		}
		for i := range moons {
			for j := range moons {
				if i >= j {
					continue
				}
				pi, pj := moons[i].pos, moons[j].pos
				dx := compare(pi.x, pj.x)
				dy := compare(pi.y, pj.y)
				dz := compare(pi.z, pj.z)
				moons[i].vel.x -= dx
				moons[j].vel.x += dx
				moons[i].vel.y -= dy
				moons[j].vel.y += dy
				moons[i].vel.z -= dz
				moons[j].vel.z += dz
			}
		}

		for i, m := range moons {
			moons[i].pos = add(m.pos, m.vel)
		}

		prev = append(prev, [3][4]int{
			{moons[0].pos.x, moons[1].pos.x, moons[2].pos.x, moons[3].pos.x},
			{moons[0].pos.y, moons[1].pos.y, moons[2].pos.y, moons[3].pos.y},
			{moons[0].pos.z, moons[1].pos.z, moons[2].pos.z, moons[3].pos.z},
		})
	}

	N := len(prev) - 1
	var period [3]int
	for i := range period {
	hunt:
		for j := 0; ; {
			// Clowning
			for j++; prev[N][i] != prev[N-j][i]; j++ {
			}
			for k := 0; k < j; k++ {
				if prev[N-j-k][i] != prev[N-k][i] {
					continue hunt
				}
			}
			period[i] = j
			break
		}
	}

	fmt.Printf("b) %d\n", lcm(lcm(period[0], period[1]), period[2]))
}

func add(u, v v3) v3 {
	return v3{u.x + v.x, u.y + v.y, u.z + v.z}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func absSum(u v3) int {
	return abs(u.x) + abs(u.y) + abs(u.z)
}

func compare(a, b int) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	default:
		return 1
	}
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
