package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type persona struct {
	Nombre, Apellido string
	Edad             int
}

func main() {
	p1 := persona{
		Nombre:   "Tomás",
		Apellido: "Cabello",
		Edad:     25,
	}

	p2 := persona{
		Nombre:   "Jaime",
		Apellido: "Garrido",
		Edad:     34,
	}

	personas := []persona{p1, p2}
	fmt.Println(personas)

	b, err := json.Marshal(personas)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
	os.Stdout.Write(b)
}
