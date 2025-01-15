package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: "abcde fghij", want: 1},
		{in: "abcde xyz ecdab", want: 0},
		{in: "a ab abc abd abf abj", want: 1},
		{in: "iiii oiii ooii oooi oooo", want: 1},
		{in: "oiii ioii iioi iiio", want: 0},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
