package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: `qjhvhtzxzqqjkmpb`, want: 1},
		{in: `xxyxx`, want: 1},
		{in: `uurcxstgmygtbstg`, want: 0},
		{in: `ieodomkazucvgmuy`, want: 0},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
