package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("El valor de x es 40")
	} else if x == 41{
		fmt.Println("El valor de x es 41")
	} else {
		fmt.Println("EL valor de x no es ni 40 ni 41")
	}
}