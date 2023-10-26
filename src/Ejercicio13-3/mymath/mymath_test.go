package mymath

import (
	"fmt"
	"testing"
)

type pairSet struct {
	values    []int
	resultado float64
}

var testsAvg = []pairSet{
	{[]int{1, 2, 3, 4, 5, 6}, 3.5},
	{[]int{10, 20, 30, 40, 50, 60}, 35},
	{[]int{0, 2, 400, 500}, 201},
	{[]int{-2, 0, 20, 40, 50}, 20},
	{[]int{90, 90, 50, 40, 100, 101}, 82.5},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range testsAvg {
		calc := CenteredAvg(v.values)
		if v.resultado != calc {
			t.Error("Expected", v.resultado, "Got", calc)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5, 6}))
	// Output:
	// 3.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6})
	}
}
