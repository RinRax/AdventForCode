package main

import (
	utils "AdventsForCode/Utils"
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
	var seen = map[point]bool{{0, 0}: true}

	for i := 1; i < data; i++{
		dir = p.move(dir, seen)
	}

	fmt.Println(p)
	return utils.Abs(p.x) + utils.Abs(p.y)
}

func (p *point) move(dir int, seen map[point]bool) int {

	if !seen[point{p.x + dirs[(dir+1)%len(dirs)].x, p.y + dirs[(dir+1)%len(dirs)].y}] {
		dir = (dir + 1) % len(dirs)
	}

	*p = point{p.x + dirs[dir].x, p.y + dirs[dir].y}
	seen[point{p.x, p.y}] = true

	return dir
}

var dirs = []point{
	{0, -1}, //N
	{1, 0},  //E
	{0, 1},  //S
	{-1, 0}, //W
}

const in = 265149
