package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Tomás")
	bar("Tomás")
	s1 := woo("Tomás")
	fmt.Println(s1)
	x, y := saludar("Tomás", "Cabello")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("Hola foo.")
}

func bar(s string) {
	fmt.Println("Hola,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo,", st)
}

func saludar(n string, a string) (string, bool){
	p := fmt.Sprint(n, " ", a, ` dice "hola".`)
	q := true
	return p, q
}