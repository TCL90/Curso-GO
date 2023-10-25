package maxMin

import (
	"testing"
)

type testpair struct {
	values []int
	max    int
}

var maxTests = []testpair{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
	{[]int{42, 56, 87, 12, 43}, 87},
	{[]int{5, 9, 32, 12, 65, 45}, 65},
	{[]int{90, 43, 65, 21, 67, 34, 23, 75}, 90},
	{[]int{90, 89, 78, 67, 56, 56, 43, 21, 12}, 90},
	{[]int{}, -1},
}

var minTests = []testpair{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1},
	{[]int{42, 56, 87, 12, 43}, 12},
	{[]int{5, 9, 32, 12, 65, 45}, 5},
	{[]int{90, 43, 65, 21, 67, 34, 23, 75}, 21},
	{[]int{90, 89, 78, 67, 56, 56, 43, 21, 12}, 12},
	{[]int{}, -1},
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error("For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.max {
			t.Error("For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
