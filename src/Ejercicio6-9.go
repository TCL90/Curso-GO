package main

import (
	"fmt"
)

func main() {
	sliceEnteros := []int{5, 6, 7, 8, 9}

	g := func(sliceE []int) int {
		if len(sliceE) == 0 {
			return 0
		}
		if len(sliceE) == 1 {
			return sliceE[0]
		}
		return sliceE[0] + sliceE[len(sliceE)-1]
	}

	fmt.Println(foo(g, sliceEnteros))
}

func foo(f func(sliceE []int) int, sliceE []int) int {
	return f(sliceE) + 1
}
