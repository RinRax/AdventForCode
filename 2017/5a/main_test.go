package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: `0
3
0
1
-3`, want: 5},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
