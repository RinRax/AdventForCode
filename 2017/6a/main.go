package main

import (
	utils "AdventsForCode/Utils"
	"fmt"
)

func main() {
	fmt.Println(solution(in))
}

func solution(data string) int {
	var seen = map[string]int{}
	var nums = utils.ToIntArray(data, "	")

	seen[utils.ToString(nums)] = 0

	for i := 1; ; i++ {
		redistribute(nums, findBiggest(nums))

		if x, ok := seen[utils.ToString(nums)]; ok {
			return i - x
		}

		seen[utils.ToString(nums)] = i
	}
}

func findBiggest(nums []int) int {
	var iMax int

	for i, v := range nums {
		if v > nums[iMax] {
			iMax = i
		}
	}

	return iMax
}

func redistribute(nums []int, i int) {
	max := nums[i]
	nums[i] = 0

	for j := range max {
		nums[(i+j+1)%len(nums)]++
	}
}

const in = `4	10	4	1	8	4	9	14	5	1	14	15	0	15	3	5`
