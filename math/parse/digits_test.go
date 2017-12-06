package parse_test

import (
	"reflect"
	"testing"

	"github.com/mrap/goutil/math/parse"
	"github.com/mrap/goutil/pkgutil"
)

func TestIntDigits(t *testing.T) {
	cases := []struct {
		in  int
		out []int
	}{
		{0, []int{0}},
		{1, []int{1}},
		{12, []int{1, 2}},
		{123, []int{1, 2, 3}},
		{1234, []int{1, 2, 3, 4}},
		{12345, []int{1, 2, 3, 4, 5}},
		{-1, []int{1}},
		{-12, []int{1, 2}},
		{-123, []int{1, 2, 3}},
		{-1234, []int{1, 2, 3, 4}},
		{-12345, []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		result := parse.IntDigits(c.in)
		if !reflect.DeepEqual(result, c.out) {
			t.Errorf(pkgutil.TestResult(result, c.out, "IntDigits(%d)", c.in))
		}
	}
}
