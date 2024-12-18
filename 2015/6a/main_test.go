package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: `turn on 0,0 through 999,999`, want: 1000000},
		{in: `toggle 0,0 through 999,0`, want: 1000},
		{in: `turn off 499,499 through 500,500`, want: 0},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
