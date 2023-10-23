package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	
	y := 42
	z := 7
	fmt.Println(y == z)
	fmt.Println(y <= z)
	fmt.Println(y >= z)
	fmt.Println(y != z)
	fmt.Println(y > z)
	fmt.Println(y < z)
}