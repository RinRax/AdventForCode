package main

import (
	utils "AdventsForCode/Utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(solution(in))
}

type point struct {
	x, y int
}

func solution(data string) int {
	var (
		points = convert(data)
		grid   = map[point]int{}
	)

	min, max := minMax(points)

	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			if res := minDist(point{x, y}, points); res >= 0 {
				grid[point{x, y}] = res
			}
		}
	}

	return maxArea(grid, points, findInfs(grid, min, max))
}

func maxArea(grid map[point]int, points []point, infs map[int]bool) int {
	var max int

	for i := range points {
		var sum int
		if infs[i] {
			continue
		}

		for _, v := range grid {
			if i == v {
				sum++
			}
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

func findInfs(grid map[point]int, min, max point) map[int]bool {
	var infs = map[int]bool{}

	for x := min.x; x <= max.x; x++ {
		infs[grid[point{x, min.y}]] = true
		infs[grid[point{x, max.y}]] = true
	}

	for y := min.y; y <= max.y; y++ {
		infs[grid[point{min.x, y}]] = true
		infs[grid[point{max.x, y}]] = true
	}

	return infs
}

func minDist(p point, points []point) int {
	var (
		count      = map[int]int{}
		min, index int
	)

	min = math.MaxInt32

	for i, point := range points {
		dist := utils.Abs(p.x-point.x) + utils.Abs(p.y-point.y)

		if dist < min {
			min = dist
			index = i
		}

		count[dist]++
	}

	if count[min] == 1 {
		return index
	}

	return -1
}

func minMax(points []point) (min, max point) {
	min = points[0]

	for _, point := range points {
		if point.x > max.x {
			max.x = point.x
		}
		if point.x < min.x {
			min.x = point.x

		}
		if point.y > max.y {
			max.y = point.y
		}
		if point.y < min.y {
			min.y = point.y
		}
	}
	max.x++
	min.x--
	max.y++
	min.y--

	return min, max
}

func convert(data string) []point {
	var points []point

	for _, line := range strings.Split(data, "\n") {
		coords := utils.StringToIntArray(strings.Split(line, ", "))
		points = append(points, point{coords[0], coords[1]})
	}

	return points
}

const in = `337, 150
198, 248
335, 161
111, 138
109, 48
261, 155
245, 130
346, 43
355, 59
53, 309
59, 189
325, 197
93, 84
194, 315
71, 241
193, 81
166, 187
208, 95
45, 147
318, 222
338, 354
293, 242
240, 105
284, 62
46, 103
59, 259
279, 205
57, 102
77, 72
227, 194
284, 279
300, 45
168, 42
302, 99
338, 148
300, 316
296, 229
293, 359
175, 208
86, 147
91, 261
188, 155
257, 292
268, 215
257, 288
165, 333
131, 322
264, 313
236, 130
98, 60`
