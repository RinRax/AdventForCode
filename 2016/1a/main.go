package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution(in1))
	fmt.Println(solution(in2))
	fmt.Println(solution(in3))
	fmt.Println(solution(in))
}

const (
	n = 1
	e = 2
	s = 3
	w = 4
)

func solution(data string) int {
	var (
		dir, x, y int
	)
	for _, v := range strings.Split(data, ", ") {
		turn := v[0]
		steps, _ := strconv.Atoi(v[1:])

		if turn == 'R' {
			dir++
			if dir > w {
				dir = n
			}
		} else {
			dir--
			if dir < n {
				dir = w
			}
		}

		if dir == n {
			y += steps
		} else if dir == e {
			x += steps
		} else if dir == s {
			y -= steps
		} else {
			x -= steps
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

const (
	in1 = `R2, L3`
	in2 = `R2, R2, R2`
	in3 = `R5, L5, R5, R3`
	in  = `R3, L5, R1, R2, L5, R2, R3, L2, L5, R5, L4, L3, R5, L1, R3, R4, R1, L3, R3, L2, L5, L2, R4, R5, R5, L4, L3, L3, R4, R4, R5, L5, L3, R2, R2, L3, L4, L5, R1, R3, L3, R2, L3, R5, L194, L2, L5, R2, R1, R1, L1, L5, L4, R4, R2, R2, L4, L1, R2, R53, R3, L5, R72, R2, L5, R3, L4, R187, L4, L5, L2, R1, R3, R5, L4, L4, R2, R5, L5, L4, L3, R5, L2, R1, R1, R4, L1, R2, L3, R5, L4, R2, L3, R1, L4, R4, L1, L2, R3, L1, L1, R4, R3, L4, R2, R5, L2, L3, L3, L1, R3, R5, R2, R3, R1, R2, L1, L4, L5, L2, R4, R5, L2, R4, R4, L3, R2, R1, L4, R3, L3, L4, L3, L1, R3, L2, R2, L4, L4, L5, R3, R5, R3, L2, R5, L2, L1, L5, L1, R2, R4, L5, R2, L4, L5, L4, L5, L2, L5, L4, R5, R3, R2, R2, L3, R3, L2, L5`
)
