package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	x = append(x[:6], x[9:]...)
	fmt.Println(x)
}
