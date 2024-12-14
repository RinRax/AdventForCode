package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{in: `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`, want: "fgij"},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %s, want: %s", got, test.want)
			}
		})

	}
}
