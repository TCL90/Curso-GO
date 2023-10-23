package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(4))
	fmt.Println(cicloFact(4))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)

}

func cicloFact(x int) int {
	factorial := 1
	for i := x; i > 0; i-- {
		factorial *= i
	}
	return factorial
}
