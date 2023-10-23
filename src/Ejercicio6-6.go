package main

import (
	"fmt"
)

func main() {
	func () {
		fmt.Println("Soy una función anónima.")
	}()

	func () {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}()
}