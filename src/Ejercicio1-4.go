package main

import "fmt"

type cuenta int
var x cuenta = 3200

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}