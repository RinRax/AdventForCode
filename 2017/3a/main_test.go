package main

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   int
		want int
	}{
		{in: 1, want: 0},
		{in: 12, want: 3},
		{in: 23, want: 2},
		{in: 1024, want: 31},
	} {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
