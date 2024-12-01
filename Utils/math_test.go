package utils_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	utils "AdventsForCode/Utils"
)

func Test_ToInt(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		in   string
		want int
	}{
		{in: "", want: 0},
		{in: "abc", want: 0},
		{in: "95", want: 95},
		{in: "0", want: 0},
		{in: "-12957", want: -12957},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()
			if got := utils.ToInt(test.in); got != test.want {
				t.Errorf("ToInt() = %d , want %d", got, test.want)
			}
		})
	}
}

func Test_ToIntArray(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		in   string
		sep  string
		want []int
	}{
		{in: "1, 2, 889, -3", sep: ", ", want: []int{1, 2, 889, -3}},
	} {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()
			if diff := cmp.Diff(test.want, utils.ToIntArray(test.in, test.sep)); diff != "" {
				t.Errorf("diff: %s", diff)
			}
		})
	}
}
