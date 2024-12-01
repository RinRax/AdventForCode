package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: "two1nine", want: 29},
		{in: "eightwothree", want: 83},
		{in: "abcone2threexyz", want: 13},
		{in: "xtwone3four", want: 24},
		{in: "4nineeightseven2", want: 42},
		{in: "zoneight234", want: 14},
		{in: "7pqrstsixteen", want: 76},
		{in: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, want: 281},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
