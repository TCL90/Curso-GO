package main

import (
	"fmt"
)

func main() {
	p := struct {
		nombre, apellido string
		edad             int
	}{
		nombre:   "Tomás",
		apellido: "Cabello",
		edad:     25,
	}

	fmt.Println(p)
}
