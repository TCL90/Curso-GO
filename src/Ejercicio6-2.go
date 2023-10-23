package main

import (
	"fmt"
)

func main() {
	sliceEnteros := []int{2, 45, 63, 98, 1}
	fmt.Println(foo(sliceEnteros...))
	fmt.Println(bar(sliceEnteros))
}

func foo(xi ...int) int {
	suma := 0
	for _, v := range xi {
		suma += v
	}
	return suma
}

func bar(xi []int) int {
	suma := 0
	for _, v := range xi {
		suma += v
	}
	return suma
}
