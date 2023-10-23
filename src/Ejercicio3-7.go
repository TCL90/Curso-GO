package main

import (
	"fmt"
)

func main() {
	var x int = 2
	if x == 2 {
		fmt.Println("x es igual a 2")
	} else if x == 3 {
		fmt.Println("x es igual a 3")
	}else {
		fmt.Println("x no es ni 2 ni 3")
	}
}