package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("La función anónima se ejecutó.")
	}()

	func(x int) {
		fmt.Println("El número es", x)
	}(2)
}

func foo() {
	fmt.Println("Foo se ejecutó.")
}
