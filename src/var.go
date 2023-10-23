package main

import "fmt"

var z = 31

func main() {
	var y int
	fmt.Println(z)
	fmt.Println(y)
	n(y)
}

func n(y int) {
	fmt.Println(z)
	fmt.Println(y)
}
