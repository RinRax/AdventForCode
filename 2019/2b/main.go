package main

import (
	utils "AdventsForCode/Utils"
	"fmt"
)

func main() {
	fmt.Println(solution(in))
}

func solution(data string) int {
	var nums = utils.ToIntArray(data, ",")
	var target = 19690720

	for i := range 99 {
		nums[1] = i

		for j := range 99 {
			nums[2] = j

		LOOP:
			for k := 0; k < len(nums)-4; k += 4 {
				if nums[k] > len(nums) || nums[k+1] > len(nums) || nums[k+2] > len(nums) || nums[k+3] > len(nums) {
					fmt.Println(k, nums[k], nums[k+1], nums[k+2], nums[k+3])
					break LOOP
				}
				switch nums[k] {
				case 1:
					nums[nums[k+3]] = nums[nums[k+1]] + nums[nums[k+2]]
				case 2:
					nums[nums[k+3]] = nums[nums[k+1]] * nums[nums[k+2]]
				case 99:
					break LOOP
				}

				if nums[0] == target {
					return 100*nums[i] + nums[j]
				}
			}

		}
	}
	return 0
}

// if something is out of range take next nr

const in = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,5,19,23,2,10,23,27,1,27,5,31,2,9,31,35,1,35,5,39,2,6,39,43,1,43,5,47,2,47,10,51,2,51,6,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,1,71,5,75,1,13,75,79,1,6,79,83,2,83,13,87,1,87,6,91,1,10,91,95,1,95,9,99,2,99,13,103,1,103,6,107,2,107,6,111,1,111,2,115,1,115,13,0,99,2,0,14,0`
