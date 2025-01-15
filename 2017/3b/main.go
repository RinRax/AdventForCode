package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution(in))
}

type point struct {
	x, y int
}

func solution(data int) int {
	var dir int
	var p point
	var seen = map[point]int{{0, 0}: 1}

	for i := 1; i < data; i++ {
		dir = p.move(dir, seen)
		if seen[p] > data {
			return seen[p]
		}
	}
	return 0
}

func (p *point) move(dir int, seen map[point]int) int {

	if seen[point{p.x + dirs[(dir+1)%len(dirs)].x, p.y + dirs[(dir+1)%len(dirs)].y}] == 0 {
		dir = (dir + 1) % len(dirs)
	}

	*p = point{p.x + dirs[dir].x, p.y + dirs[dir].y}
	seen[point{p.x, p.y}] = seen[point{p.x, p.y + 1}] +
		seen[point{p.x, p.y - 1}] +
		seen[point{p.x + 1, p.y}] +
		seen[point{p.x + 1, p.y + 1}] +
		seen[point{p.x + 1, p.y - 1}] +
		seen[point{p.x - 1, p.y}] +
		seen[point{p.x - 1, p.y + 1}] +
		seen[point{p.x - 1, p.y - 1}]

	return dir
}

var dirs = []point{
	{0, -1}, //N
	{1, 0},  //E
	{0, 1},  //S
	{-1, 0}, //W
}

const in = 265149
