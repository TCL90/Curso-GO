package main

import (
	"fmt"
	"perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Black",
		edad:   perro.EdadPerro(5),
	}
	fmt.Println(c1)
}
