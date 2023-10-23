package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[:])
	fmt.Println(x[1:3])

	for i := 0; i <= len(x); i++ {
		fmt.Println(i, " ", x[i])
	}
}