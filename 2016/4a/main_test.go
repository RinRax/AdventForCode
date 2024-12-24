package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{in: `aaaaa-bbb-z-y-x-123[abxyz]`, want: 123},
		{in: `a-b-c-d-e-f-g-h-987[abcde]`, want: 987},
		{in: `not-a-real-room-404[oarel]`, want: 404},
		{in: `totally-real-room-200[decoy]`, want: 0},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}
