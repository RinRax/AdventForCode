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
		{in: 800, want: 806},
	} {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
