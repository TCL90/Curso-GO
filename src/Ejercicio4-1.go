package main

import (
	"fmt"
)

func main() {
	var array [5]int
	array[0] = 21
	array[1] = 22
	array[2] = 23
	array[3] = 24
	array[4] = 25

	var array = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T\n\n", array)
	for i, v := range array {
		fmt.Println(i, " ", v)
	}
}
