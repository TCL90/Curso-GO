package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println()

	for i, v := range x {
		fmt.Println(i, " ", v)
	}

	y := []int{2, 4, 6, 8, 10}
	fmt.Println(y)
	fmt.Printf("%T", y)
}