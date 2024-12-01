package main

import (
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct{
		in string
		want int
	}{
		{in: "", want: 0},
		{in: "1abc2", want: 12},
		{in: "pqr3stu8vwx", want: 38},
		{in: "a1b2c3d4e5f", want: 15},
		{in: "treb7uchet", want: 77},
	}{
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()
			
			if got := solution(test.in); got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})

	}
}