package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Eduar",
		Edad:   32,
	}

	u2 := usuario{
		Nombre: "Juan",
		Edad:   27,
	}

	u3 := usuario{
		Nombre: "Che",
		Edad:   54,
	}

	usuarios := []usuario{u1, u2, u3}

	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(usuarios)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
