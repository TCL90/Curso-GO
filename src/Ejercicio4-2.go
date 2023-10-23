package main

import (
	"fmt"
)

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("%T\n\n", slice)
	for i, v := range slice {
		fmt.Println(i, " ", v)
	}
}
