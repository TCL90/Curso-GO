package main

import (
	"fmt"
)

func main() {
	var a string = `Esto es un raw string literal,
	es decir,
	no,
	interpretado.`
	fmt.Println(a)
}