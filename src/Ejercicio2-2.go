package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 3

	c := (a == b)
	d := (a <= b)
	e := (a >= b)
	f := (a != b)
	g := (a < b)
	h := (a > b)

	fmt.Printf("Para los valores %d y %d\n", a, b)
	fmt.Printf("¿Son iguales?: %t\n", c)
	fmt.Printf("¿%d es menor o igual que %d?: %t\n", a, b, d)
	fmt.Printf("¿%d es mayor o igual que %d?: %t\n", a, b, e)
	fmt.Printf("¿%d es distinto de %d?: %t\n", a, b, f)
	fmt.Printf("¿%d es menor que %d?: %t\n", a, b, g)
	fmt.Printf("¿%d es mayor que %d?: %t\n", a, b, h)
}
