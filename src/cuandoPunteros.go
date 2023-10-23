package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x antes de foo", x)
	foo(x)
	fmt.Println("x después de foo", x)

	fmt.Println()

	fmt.Println("x antes de foo2", x)
	foo2(&x)
	fmt.Println("x después de foo2", x)
}

func foo(y int) {
	fmt.Println("y antes de cambiar su valor en foo", y)
	y = 43
	fmt.Println("y después de cambiar su valor en foo", y)
}

func foo2(y *int) {
	fmt.Println("y antes de cambiar su valor en foo", *y)
	*y = 43
	fmt.Println("y después de cambiar su valor en foo", *y)
}
