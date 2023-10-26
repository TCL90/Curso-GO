package main

import (
	"average"
	"fmt"
)

func main() {
	fmt.Println("La media de 1 y 2 es:", average.Media([]int{1, 2}...))
	fmt.Println("La media de 1, 2, 3, 4, 5, 6, 7, 8 es:", average.Media([]int{1, 2, 3, 4, 5, 6, 7, 8}...))
	fmt.Println("La media de 11, 12 es:", average.Media([]int{11, 12}...))
	fmt.Println("La media de 34, 25 es:", average.Media([]int{34, 25}...))
	fmt.Println("La media de 14, 56, 85, 32, 12 es:", average.Media([]int{14, 56, 85, 32, 12}...))
}
