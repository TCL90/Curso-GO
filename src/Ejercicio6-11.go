package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(4))
}

func fibonacci(x int) int {
	if x == 1 {
		return 1
	}
	if x == 2 {
		return 1
	}

	return fibonacci(x-1) + fibonacci(x-2)

}
