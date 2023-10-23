package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type contacto struct {
		Nombre    string
		Telefono  int
		Etiquetas []string
	}

	c1 := contacto{
		Nombre:    "Tomás",
		Telefono:  658310185,
		Etiquetas: []string{"Personal", "Nuevo"},
	}

	fmt.Println(c1)
	bs, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
	os.Stdout.Write(bs)
}
