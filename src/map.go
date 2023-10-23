package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin": 27
	}

	fmt.Println(m)
	fmt.Println(m["Batman"])
	fmt.Println(m["Robin"])

	v, ok := m["Tomás"]
	fmt.Println(v, ok)

	if v, ok := m["Tomás"]; ok {
		fmt.Println("El valor de la clave Tomás es:", v)
	}
}