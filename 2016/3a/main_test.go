package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: ` 5     10  25`, want: 0},
		{in: `    15 10       25`, want: 1},
		{in: ` 2900     2800  99`, want: 0},
		{in: `        1 5 3`, want: 0},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
