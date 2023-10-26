package mate

import (
	"fmt"
	"testing"
)

type pairSet struct {
	values    []int
	respuesta int
}

var testsMate []pairSet = []pairSet{
	{[]int{12, 13}, 25},
	{[]int{12, 13, 45}, 70},
	{[]int{12, 13, 23, 23}, 71},
	{[]int{12, 90, 23}, 125},
	{[]int{1, 2, 3, 4}, 10},
}

func TestSuma(t *testing.T) {
	for _, v := range testsMate {
		got := Sum(v.values...)
		if got != v.respuesta {
			t.Errorf("Expected %d, got %d for values %v", v.respuesta, got, v.values)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(2, 4, 5))
	// Output:
	// 11
}
