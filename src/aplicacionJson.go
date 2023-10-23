package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[
		{"Nombre":"Tomás","Apellido":"Cabello","Edad":25},
		{"Nombre":"Jaime","Apellido":"Garrido","Edad":34}
		]`
	bs := []byte(s)

	fmt.Printf("%T\n%T\n", s, bs)

	var personas []persona

	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(personas)
	fmt.Println()

	for i, v := range personas {
		fmt.Println("\nPersona número", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}