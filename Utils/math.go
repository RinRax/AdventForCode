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

	for i := range str {
		res[i] = ToInt(str[i])
	}
	return res
}

func StringToIntArray(s []string) []int {
	var res []int

	for _, v := range s {
		res = append(res, ToInt(v))
	}

	return res
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func IsEven(i int) bool {
	return i == i/2*2
}

func ToString(ints []int) string {
	var res string

	for _, v := range ints {
		res += strconv.Itoa(v) + " "
	}

	return res
}
