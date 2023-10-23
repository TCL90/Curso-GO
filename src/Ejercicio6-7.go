package main

import (
	"fmt"
)

var x int
var g func() = func() {
	fmt.Println("g desde afuera.")
}

func main() {
	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	g()
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", g)
}
