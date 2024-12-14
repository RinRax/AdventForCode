package main

import (
	utils "AdventsForCode/Utils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solution(in))
}

func solution(data string) int {
	var stones = utils.ToIntArray(data, " ")

	for b := 0; b < 25; b++ {
		length := len(stones)

		for i := 0; i < length; i++ {
			newStones := check(stones[i])

			switch len(newStones) {
			case 1:
				stones[i] = newStones[0]
			case 2:
				stones[i] = newStones[0]
				stones = append(stones, newStones[1])
			}
		}

		fmt.Println(len(stones))
	}
	return len(stones)
}

func check(stone int) []int {
	switch {
	case stone == 0:
		return []int{1}
	case utils.IsEven(len(strconv.Itoa(stone))):
		return splitStone(stone)
	default:
		return []int{stone * 2024}
	}
}

func splitStone(stone int) []int {
	s := strconv.Itoa(stone)

	return []int{utils.ToInt(s[:len(s)/2]), utils.ToInt(s[len(s)/2:])}
}

// for each blink, go through each stone
// check all rules and return array of new stones (either 1 or 2 items)

const in = `64554 35 906 6 6960985 5755 975820 0`
