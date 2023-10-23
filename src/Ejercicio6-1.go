package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 2
}

func bar() (int, string) {
	return 2, "Hola"
}
