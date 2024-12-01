package utils

import (
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToIntArray(s string, sep string) []int {
	str := strings.Split(s, sep)
	res := make([]int, len(str))

	for i, v := range str {
		res[i] = ToInt(v)
	}
	return res
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
