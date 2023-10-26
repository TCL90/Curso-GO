package average

import "testing"

type pairSet struct {
	values    []int
	resultado float64
}

var testMedias = []pairSet{
	{[]int{1, 2}, 1.5},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4.5},
	{[]int{11, 12}, 11.5},
	{[]int{34, 25}, 29.5},
	{[]int{14, 56, 85, 32, 12}, 39.8},
}

func TestMedia(t *testing.T) {
	for _, v := range testMedias {
		got := Media(v.values...)
		if got != v.resultado {
			t.Errorf("Expected %f, Got %f from Average of %v", v.resultado, got, v.values)
		}
	}
}
