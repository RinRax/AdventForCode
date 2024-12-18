package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: `ugknbfddgicrmopn`, want: 1},
		{in: `aaa`, want: 1},
		{in: `jchzalrnumimnmhp`, want: 0},
		{in: `haegwjzuvuyypxyu`, want: 0},
		{in: `dvszwmarrgswjxmb`, want: 0},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
