package slices_test

import (
	"reflect"
	"testing"

	"github.com/mrap/goutil/pkgutil"
	"github.com/mrap/goutil/slices"
)

func TestIntsReverse(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{[]int{0}, []int{0}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, c := range cases {
		in := c.in[:]
		slices.ReverseInts(in)
		if !reflect.DeepEqual(in, c.out) {
			t.Errorf(pkgutil.TestResult(in, c.out, "ReverseInts(%s)", c.in))
		}
	}
}
