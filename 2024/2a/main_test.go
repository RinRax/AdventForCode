package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: `7 6 4 2 1`, want: 1},
		{in: `1 2 7 8 9`, want: 0},
		{in: `9 7 6 2 1`, want: 0},
		{in: `1 3 2 4 5`, want: 0},
		{in: `8 6 4 4 1`, want: 0},
		{in: `1 3 6 7 9`, want: 1},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
